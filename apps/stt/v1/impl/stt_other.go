package impl

// "github.com/gordonklaus/portaudio"

// 捕获麦克风音频流
// func captureAudio(audioStream chan<- []byte) {
// 	// 设置音频流参数
// 	inputBuffer := make([]int16, 1024)
// 	stream, err := portaudio.OpenDefaultStream(1, 0, 16000, len(inputBuffer), inputBuffer)
// 	if err != nil {
// 		exception.NewInternalServerError("error: %v", err)
// 	}
// 	defer stream.Close()

// 	// 开始音频流
// 	if err := stream.Start(); err != nil {
// 		exception.NewInternalServerError("error: %v", err)
// 	}
// 	defer stream.Stop()

// 	// 将音频数据传输到通道
// 	for {
// 		if err := stream.Read(); err != nil {
// 			exception.NewInternalServerError("error: %v", err)
// 		}
// 		// 将捕获到的音频数据转换为字节切片并发送到音频流通道
// 		audioBytes := make([]byte, len(inputBuffer)*2)
// 		for i, sample := range inputBuffer {
// 			audioBytes[i*2] = byte(sample)
// 			audioBytes[i*2+1] = byte(sample >> 8)
// 		}
// 		audioStream <- audioBytes
// 	}
// }
