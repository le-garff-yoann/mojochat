<template>
  <b-container id="app" fluid>
    <b-row>
      <b-col>
        <b-form-input
          id="bodyInput"
          placeholder="Write a message..."
          v-model="bodyInput"
          @keydown.enter.native="submitMessageBody"

          class="bubble"
          v-bind:disabled="! ws.isConnected"
        >
        </b-form-input>
      </b-col>
    </b-row>

    <b-row>
      <b-col id="blabla">
        <b-container
          v-for="(msg, i) in messages"
          :key="i"

          fluid
          class="bubble"
          v-bind:class="msg.me ? 'me' : 'others'"
        >
          <p class="messageBody">{{ msg.body }}</p><span class="meta-right"><span class="messageUUID">{{ msg.uuid }}</span> @ <span class="messageDatetime">{{ msg.datetime }}</span></span>
        </b-container>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
import Vue from 'vue'
import VueNativeSock from 'vue-native-websocket'
Vue.use(VueNativeSock, 'ws://localhost:8080', { connectManually: true })

import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
Vue.use(BootstrapVue);

export default {
  name: 'app',
  data() {
    return {
      bodyInput: '',
      messages: [],
      ws: { isConnected: false }
    }
  },
  created() {
    this.$options.sockets.onopen = () => this.ws.isConnected = true
    this.$options.sockets.onclose = () => this.ws.isConnected = false

    this.$options.sockets.onmessage = (e) => this.messages.unshift(JSON.parse(e.data))

    setInterval(() => {
      this.ws.isConnected || this.$connect(`ws${window.location.protocol === 'https:' ? 's' : ''}://${window.location.host}${process.env.VUE_APP_ROOT_API}/chat`)
    }, 1000)
  },
  methods: {
    submitMessageBody(e) {
      e.preventDefault()

      try {
        this.$socket.send(this.bodyInput)
      } catch (ex) {
        throw ex
      }

      this.bodyInput = null
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
