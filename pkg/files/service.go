package files

import (
	"bufio"
	"context"
	"encoding/binary"
	"file-editor/api"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Service struct {
	Producer   chan []byte
	Consumer   chan []byte
	waitGroups map[string]*sync.WaitGroup
}

func New(ch chan []byte) *Service {
	service := &Service{
		Producer:   ch,
		Consumer:   ch,
		waitGroups: make(map[string]*sync.WaitGroup),
	}

	go service.RunConsumer()

	return service
}

func (s *Service) ReadFile(ctx context.Context, fileName string) (*api.ReadFileResponse, error) {
	// Obtenenemos la ruta del archivo
	path, err := s.getFilePath(fileName)
	if err != nil {
		return nil, err
	}

	// Leemos el archivo
	content, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, NewFileNotFoundError("solicitud inválida, el archivo no existe")
		}

		return nil, err
	}

	return &api.ReadFileResponse{Content: string(content)}, nil
}

// SaveFile crea un nuevo archivo usando el contenido proporcionado en la solicitud
func (s *Service) SaveFile(ctx context.Context, request *api.SaveFileRequest) (*api.SaveFileResponse, error) {
	fileBytes := request.Content
	fileName := request.Filename

	// Creamos un header que indique la longitud del nombre de archivo
	header := make([]byte, headerSize)
	header[lastMessageHeaderIndex] = byte(notLastMessage)
	binary.LittleEndian.PutUint32(header[fileNameHeaderStartIndex:fileNameHeaderEndIndex], uint32(len(fileName)))

	// Obtenemos la ruta del archivo donde se almacenara
	path, err := s.getFilePath(fileName)
	if err != nil {
		return nil, err
	}

	// Creamos el nuevo archivo, sobrescribiéndolo si ya existe
	_, err = os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("no se pudo crear un nuevo archivo: %w", err)
	}

	// Creamos un WaitGroup para sincronizar la escritura del archivo
	wg := &sync.WaitGroup{}
	wg.Add(1)
	s.waitGroups[fileName] = wg

	// Iteramos sobre los bytes del archivo y los enviamos en el channel
	for i := 0; i < len(fileBytes); i += bytesSize {
		select {
		case <-ctx.Done():
			// Si el contexto se cancela, retornar un error indicando la cancelación
			return nil, NewContextDoneError(fmt.Sprintf("contexto cancelado durante la escritura del archivo '%s': %s", fileName, ctx.Err().Error()))
		default:
			end := i + bytesSize

			// Actualizamos el header si es el último mensaje a enviar
			if end >= len(fileBytes) {
				header[lastMessageHeaderIndex] = byte(lastMessage)
				end = len(fileBytes)
			}

			// Armamos los bytes a enviar, incluyendo el header y los datos del archivo
			bytesToSend := append(header, append([]byte(fileName), fileBytes[i:end]...)...)

			// Enviamos los bytes en el channel
			s.Producer <- bytesToSend
		}
	}

	// Esperamos a que se complete la escritura del archivo
	wg.Wait()
	delete(s.waitGroups, fileName)

	return &api.SaveFileResponse{
		Response: fmt.Sprintf("Archivo '%s' ha sido subido exitosamente", fileName),
	}, nil
}

func (s *Service) FindText(ctx context.Context, req *api.FindTextRequest) (*api.FindTextResponse, error) {
	// Obtenemos la ruta del archivo
	path, err := s.getFilePath(req.Filename)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo la ruta del archivo: %w", err)
	}

	// Abrimos el archivo en modo lectura
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, NewFileNotFoundError("solicitud inválida, el archivo no existe")
		}

		return nil, fmt.Errorf("error abriendo el archivo: %w", err)
	}

	defer file.Close()

	// Variables para almacenar las líneas que contienen el texto y el contador de ocurrencias
	var lines []string
	var count int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, req.SearchText) {
			lines = append(lines, line)
			count += int64(strings.Count(line, req.SearchText))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error escaneando el archivo: %w", err)
	}

	// Retornamos la cantidad de ocurrencias y las líneas que coinciden
	return &api.FindTextResponse{
		Count: count,
		Lines: lines,
	}, nil
}

// ReadAllFiles lee todos los archivos en el working directory y retorna una respuesta que contiene el nombre y el contenido de cada archivo
func (s *Service) ReadAllFiles(ctx context.Context, req *api.Empty) (*api.ReadAllFilesResponse, error) {
	// return s.readAllFilesSynchronously() //Para comparar la ventaja usando threads,
	return s.readAllFilesConcurrently()
}

// RunConsumer simula un consumidor que lee bytes del channel Consumer y los escribe en archivos
func (s *Service) RunConsumer() {
	for bytes := range s.Consumer {
		if len(bytes) < headerSize {
			fmt.Println("error: slice de bytes inválido")
			continue
		}

		// Obtenemos la longitud del nombre de archivo del header
		fileNameLen := int(binary.LittleEndian.Uint32(bytes[fileNameHeaderStartIndex:fileNameHeaderEndIndex]))

		if len(bytes) < headerSize+fileNameLen {
			fmt.Println("error: slice de bytes inválido")
			continue
		}

		// Extraemos el nombre del archivo y los datos del byte slice recibido
		filename := string(bytes[headerSize : headerSize+fileNameLen])
		data := bytes[headerSize+fileNameLen:]

		// Obtenemos la ruta del archivo donde se escribirán los datos
		path, err := s.getFilePath(filename)
		if err != nil {
			fmt.Printf("error obteniendo la ruta del archivo: %s\n", err.Error())
			continue
		}

		// Abrimos el archivo en modo escritura, creándolo si no existe y agregando al final si existe
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Printf("error abriendo el archivo: %s\n", err.Error())
			continue
		}

		// Escribimos los datos en el archivo
		_, err = file.Write(data)
		if err != nil {
			fmt.Printf("error escribiendo en el archivo: %s\n", err.Error())
			file.Close()
			continue
		}

		// Cerramos el archivo después de escribir
		err = file.Close()
		if err != nil {
			fmt.Printf("error cerrando el archivo: %s\n", err.Error())
			continue
		}

		// Si es el último mensaje para este archivo, marcar la WaitGroup como completada
		if bytes[lastMessageHeaderIndex] == lastMessage {
			if wg, exists := s.waitGroups[filename]; exists {
				wg.Done()
			}
		}

		// Esperamos un segundo antes de procesar el próximo byte slice
		time.Sleep(time.Second)
	}
}

func (s *Service) readAllFilesConcurrently() (*api.ReadAllFilesResponse, error) {
	start := time.Now()

	wg := &sync.WaitGroup{}

	results := make(chan FileContent)

	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error obteniendo el directorio de trabajo actual: %w", err)
	}

	directory := fmt.Sprintf("%s/%s", cwd, filesFolder)

	// Leemos los archivos en el directorio especificado
	directories, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("no se pudo leer el directorio: %w", err)
	}

	// Iteramos sobre todos los archivos en el working directory
	for _, file := range directories {
		if !file.IsDir() {

			wg.Add(1)
			go func(file os.DirEntry) {
				// Agregamos un mini sleep para comprobar la diferencia de tiempo de ejecucion usando threads
				time.Sleep(1 * time.Second)

				defer wg.Done()

				content, err := os.ReadFile(filepath.Join(directory, file.Name()))
				if err != nil {
					log.Printf("no se pudo leer el archivo %s, skippeando: %v", file.Name(), err)
					return
				}

				results <- FileContent{
					Name:    file.Name(),
					Content: string(content),
				}
			}(file)
		}
	}

	// Esperamos a que finalice la ejecucion de todos los threads creados anteriormente, y cerramos el channel
	go func() {
		wg.Wait()
		close(results)
	}()

	response := &api.ReadAllFilesResponse{}

	// Esta es la forma mas sencilla de iterar un vector de cualquier tipo en Go, pero tambien se admite la forma -> for int i:= 0; i < len(results); i++ {}
	for result := range results {
		response.Content = append(response.Content, &api.FileContent{
			Name:    result.Name,
			Content: result.Content,
		})
	}

	// Calculamos el response time en segundos
	response.ResponseTime = float32(time.Since(start).Seconds())

	return response, nil
}

// readAllFilesSynchronously lee todos los archivos en un directorio especificado de manera sincrona y
// retorna una respuesta que contiene el nombre y el contenido de cada archivo.
func (s *Service) readAllFilesSynchronously() (*api.ReadAllFilesResponse, error) {
	start := time.Now()

	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error obteniendo el directorio de trabajo actual: %w", err)
	}

	directory := fmt.Sprintf("%s/%s", cwd, filesFolder)

	// Leemos los archivos en el directorio especificado
	directories, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("no se pudo leer el directorio: %w", err)
	}

	response := &api.ReadAllFilesResponse{}

	// Iteramos sobre todos los archivos en el working directory
	for _, file := range directories {
		if !file.IsDir() {
			// Agregamos un mini sleep para comprobar la diferencia de tiempo iterando de forma sincrona
			time.Sleep(1 * time.Second)
			content, err := os.ReadFile(filepath.Join(directory, file.Name()))
			if err != nil {
				log.Printf("no se pudo leer el archivo %s, skippeando: %v", file.Name(), err)
				continue
			}

			response.Content = append(response.Content, &api.FileContent{
				Name:    file.Name(),
				Content: string(content),
			})
		}
	}

	// Calculamos el response time en segundos
	response.ResponseTime = float32(time.Since(start).Seconds())

	return response, nil
}

// getFilePath devuelve la ruta completa del archivo en el sistema de archivos
func (s *Service) getFilePath(filename string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error obteniendo el directorio de trabajo actual: %w", err)
	}

	return filepath.Join(cwd, filesFolder, filename), nil
}
