<template>
  <div>
    <h1>实时音频录制 (PCM 格式)</h1>
    <button @click="startRecording" :disabled="isRecording">开始录音</button>
    <button @click="stopRecording" :disabled="!isRecording">停止录音</button>
    <div v-if="transcript">
      <h2>识别结果：</h2>
      <p>{{ transcript }}</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      audioContext: null, // 音频上下文
      mediaStream: null, // 麦克风音频流
      audioWorkletNode: null, // AudioWorkletNode 实例
      isRecording: false, // 录音状态
      socket: null, // WebSocket 连接
      transcript: '' // 保存后端返回的识别结果
    }
  },
  methods: {
    async startRecording() {
      try {
        // 初始化 WebSocket
        this.socket = new WebSocket('ws://localhost:8010/mcenter/api/v1/SpeechToTextV1/ws')

        this.socket.onopen = () => {
          console.log('WebSocket 连接成功')
        }

        this.socket.onerror = (error) => {
          console.error('WebSocket 错误:', error)
        }

        this.socket.onclose = () => {
          console.log('WebSocket 连接关闭')
        }

        // 接收后端返回的数据
        this.socket.onmessage = (event) => {
          const message = event.data
          console.log('收到后端数据:', message)
          this.transcript += message // 将新的识别结果追加到页面上
        }

        // 获取麦克风音频流
        this.mediaStream = await navigator.mediaDevices.getUserMedia({ audio: true })

        // 创建 AudioContext 实例
        this.audioContext = new (window.AudioContext || window.webkitAudioContext)()
        const source = this.audioContext.createMediaStreamSource(this.mediaStream)

        // 动态加载 AudioWorklet 处理器
        const workletCode = `
          class AudioProcessor extends AudioWorkletProcessor {
            process(inputs, outputs, parameters) {
              const input = inputs[0];
              if (input.length > 0) {
                const channelData = input[0];
                const pcmData = new Int16Array(channelData.length);

                // 将浮点音频数据转换为 16 位 PCM 数据
                for (let i = 0; i < channelData.length; i++) {
                  pcmData[i] = Math.max(-1, Math.min(1, channelData[i])) < 0 ?
                    Math.floor(channelData[i] * 32768) : Math.ceil(channelData[i] * 32767);
                }

                // 通过 port 发送 PCM 数据到主线程
                this.port.postMessage(pcmData.buffer);
              }
              return true;
            }
          }

          registerProcessor('audio-processor', AudioProcessor);
        `

        // 动态创建工作模块
        const workletBlob = new Blob([workletCode], { type: 'application/javascript' })
        const workletUrl = URL.createObjectURL(workletBlob)
        await this.audioContext.audioWorklet.addModule(workletUrl)

        // 创建 AudioWorkletNode 节点
        this.audioWorkletNode = new AudioWorkletNode(this.audioContext, 'audio-processor')

        // 连接音频流源和 AudioWorkletNode
        source.connect(this.audioWorkletNode)
        this.audioWorkletNode.connect(this.audioContext.destination)

        // 监听 AudioWorkletNode 处理后的 PCM 数据并通过 WebSocket 发送
        this.audioWorkletNode.port.onmessage = (event) => {
          const pcmData = event.data // 这是从 AudioWorklet 传递来的 ArrayBuffer
          if (this.socket && this.socket.readyState === WebSocket.OPEN) {
            this.socket.send(pcmData) // 直接发送 ArrayBuffer 数据
          }
        }

        this.isRecording = true // 更新录音状态
      } catch (err) {
        console.error('录音启动失败:', err)
      }
    },

    stopRecording() {
      // 停止录音并释放资源
      if (this.audioWorkletNode) {
        this.audioWorkletNode.disconnect() // 断开音频节点
      }

      if (this.audioContext) {
        this.audioContext.close() // 关闭音频上下文
      }

      if (this.mediaStream) {
        this.mediaStream.getTracks().forEach((track) => track.stop()) // 停止麦克风音频流
      }

      if (this.socket && this.socket.readyState === WebSocket.OPEN) {
        this.socket.close() // 关闭 WebSocket 连接
      }

      this.isRecording = false // 更新录音状态
    }
  }
}
</script>

<style scoped>
button {
  margin: 10px;
}
</style>
