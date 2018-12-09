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
      { datetime: '<p>0</p>', uuid: '<p>0</p>', body: '<p>0</p>', me: true },
      { datetime: '<p>1</p>', uuid: '<p>1</p>', body: '<p>1</p>', me: false }
    ]

    for (const msg of messages) {
      wrapper.vm.messages.push(msg)
    }

    const bubbles = blabla.findAll('.bubble')

    for (const [i, msg] of messages.entries()) {
      const bubble = bubbles.at(i).find('div')

      expect(bubble.classes()).toContain(msg.me ? 'me' : 'others')

      expect(bubble.find('.messageBody').text()).toBe(msg.body)
      expect(bubble.find('.messageUUID').text()).toBe(msg.uuid)
      expect(bubble.find('.messageDatetime').text()).toBe(msg.datetime)
    }
  })
})
