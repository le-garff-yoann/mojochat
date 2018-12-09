<template>
  <div id="app" class="container-fluid">
    <div class="row">
      <div class="col">
        <input
          type="text"
          placeholder="Write a message..."
          class="container-fluid bubble"
          ref="bodyInput"
          @keyup.enter="submitMessageBody"
          v-bind:disabled="! ws.isConnected"
        />
      </div>
    </div>
    <div class="row">
      <div class="col" id="blabla">
        <div
          class="container-fluid bubble"
          v-for="(msg, i) in messages"
          :key="i"
          v-bind:class="msg.me ? 'me' : 'others'"
        >
          <p class="messageBody">{{ msg.body }}</p><span class="meta-right"><span class="messageUUID">{{ msg.uuid }}</span> @ <span class="messageDatetime">{{ msg.datetime }}</span></span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import VueNativeSock from 'vue-native-websocket'

Vue.use(VueNativeSock, 'ws://localhost:8080', { connectManually: true })

export default {
  name: 'app',
  data() {
    return {
      messages: [],
      ws: { isConnected: false }
    }
  },
  created() {
    this.$options.sockets.onopen = () => this.ws.isConnected = true
    this.$options.sockets.onclose = () => this.ws.isConnected = false

    this.$options.sockets.onmessage = (e) => this.messages.unshift(JSON.parse(e.data))

    this.$connect(`ws${window.location.protocol === 'https:' ? 's' : ''}://${window.location.host}${window.location.pathname}chat`)
  },
  methods: {
    submitMessageBody(e) {
      let bodyInput = this.$refs.bodyInput

      if (bodyInput.value.length) {
        try {
          this.$socket.send(bodyInput.value)

          bodyInput.value = null
        } catch (ex) {
          throw ex
        }
      }

      e.preventDefault()
    }
  }
}
</script>

<style>
#app {
  font-family: 'Consolas', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  background-color: lightblue;
  overflow: hidden;
}

#blabla {
  display: block;
  height: 95vh;
  overflow-y: scroll;
}

.bubble {
  border: 2px solid #dedede;
  border-radius: 5px;
  padding: 10px;
  margin: 10px 0;
}

.bubble::after {
  content: '';
  clear: both;
  display: block;
}

.me {
  background-color: lightgreen;
}

.others {
  background-color: yellow;
}

.meta-right {
  float: right;
}

.messageBody {
  font-weight: normal;
}

.messageUUID {
  font-weight: bold;
}

.messageDatetime {
  font-style: italic;
}
</style>
