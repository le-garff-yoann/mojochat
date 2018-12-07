import { mount } from '@vue/test-utils' 
import App from '@/App.vue'

describe('App.vue', () => {
  it('has a data hook', () => {
    expect(typeof App.data).toBe('function')
  })

  it('has a created hook', () => {
    expect(typeof App.created).toBe('function')
  })

  const data = App.data()

  it('sets the correct default data', () => {
    expect(typeof data.messages).toBe('object')
    expect(typeof data.ws).toBe('object')
    expect(typeof data.ws.isConnected).toBe('boolean')
    expect(data.ws.isConnected).toBeFalsy()
  })

  it('has the submitMessageBody method', () => {
    expect(typeof App.methods.submitMessageBody).toBe('function')
  })

  const wrapper = mount(App)

  it('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })

  const app = wrapper.find(App)

  const bodyInput = app.find({ ref: 'bodyInput' })
  const blabla = app.find('#blabla')

  it('renders app', () => {
    expect(app.is('div')).toBe(true)

    expect(bodyInput.is('input')).toBe(true)
    expect(bodyInput.attributes('type')).toBe('text')
    expect(bodyInput.attributes('placeholder')).toBe('Write a message...')

    expect(blabla.is('div')).toBe(true)
  })

  it('renders messages', () => {
    const messages = [
      { datetime: 0, uuid: 0, body: 0, me: true },
      { datetime: 1, uuid: 1, body: 1, me: false }
    ]

    for (const msg of messages) {
      wrapper.vm.messages.push(msg)
    }

    const bubbles = blabla.findAll('.bubble')

    for (const [i, msg] of messages.entries()) {
      expect(bubbles.at(i).html()).toBe(`<div class="container-fluid bubble ${msg.me ? 'me' : 'others'}"><p>${msg.body}</p><span class="meta-right"><b>${msg.uuid}</b> @ <i>${msg.datetime}</i></span></div>`)
    }
  })
})

