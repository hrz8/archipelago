<template>
  <div class="d-flex justify-content-center align-items-center vh-100">
    <div class="card w-50">
      <div class="card-body">
        <h5 class="card-title text-center">{{ `Hey ${meKey}, ${chatTitle}` }}</h5>
        <div ref="scrollable" class="chat-window" style="height: 300px; overflow-y: auto;">
          <div v-for="(message, index) in messages" :key="index" class="mb-3 d-flex" :class="{'justify-content-end': message.sender === 'me', 'justify-content-start': message.sender !== 'me'}">
            <div :class="['chat-bubble', message.sender === 'me' ? 'bg-primary text-white' : 'bg-secondary text-white']">
              <small>{{ message.sender === 'me' ? 'You' : 'Other Person' }}:</small>
              <p>{{ message.text }}</p>
            </div>
          </div>
        </div>
        <div class="input-group mt-3">
          <input
            type="text"
            class="form-control"
            placeholder="Type a message..."
            v-model="messageText"
            @keyup.enter="sendMessage"
            :disabled="!isRoomFull"
          />
          <button class="btn btn-primary" @click="sendMessage" :disabled="!isRoomFull">Send</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, onUnmounted, nextTick } from 'vue';

export default {
  name: 'ChatWindow',
  setup() {
    const scrollable = ref(null);
    const messages = ref([]);
    const messageText = ref('');
    const chatTitle = ref('Connecting...');
    const meKey = ref(null);
    const isRoomFull = ref(false);
    let ws;

    const scrollToBottom = () => {
      if (scrollable.value) {
        scrollable.value.scrollTop = scrollable.value.scrollHeight;
      }
    }

    const initializeWebSocket = () => {
      ws = new WebSocket('ws://localhost:8080/ws');

      ws.onmessage = (event) => {
        const message = event.data;
        if (message.includes("Connected")) {
          meKey.value = message.split(' ').pop();
        } else if (message.includes("User disconnected")) {
          alert("User disconnected, please refresh.");
          window.location.reload();
        } else if (message.includes("Room is full")) {
          alert("Room is full, please try again later.");
          ws.close();
        } else if (message.includes("Chat started!")) {
          chatTitle.value = `Chat with ${message.split(' ').pop()}`;
          isRoomFull.value = true;
        } else {
          messages.value.push({ sender: "Other Person", text: message });
          nextTick(() => {
            scrollToBottom();
          });
        }
      };
    };

    const sendMessage = () => {
      if (messageText.value.trim()) {
        ws.send(messageText.value);
        messages.value.push({ sender: "me", text: messageText.value });
        messageText.value = '';
        nextTick(() => {
          scrollToBottom();
        });
      }
    };

    onMounted(() => {
      initializeWebSocket();
    });

    onUnmounted(() => {
      if (ws) ws.close();
    });

    return {
      scrollable,
      messages,
      messageText,
      chatTitle,
      meKey, 
      isRoomFull,
      sendMessage,
    };
  },
};
</script>

<style scoped>
.chat-window {
  max-height: 400px;
  overflow-y: auto;
  border: 1px solid #ddd;
  padding: 15px;
  margin-top: 15px;
}

.chat-bubble {
  padding: 10px;
  border-radius: 15px;
  margin: 5px;
  max-width: 70%;
  word-wrap: break-word;
}

.bg-primary {
  border-radius: 15px 15px 0px 15px;
}

.bg-secondary {
  border-radius: 15px 15px 15px 0px;
}
</style>
