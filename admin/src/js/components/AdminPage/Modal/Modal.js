
import React from 'react'
import ModalNewProjectContent from './ModalNewProjectContent'
import ModalNewExpeditionContent from './ModalNewExpeditionContent'

class Modal extends React.Component {
  constructor (props) {
    super(props)
    this.getContent = this.getContent.bind(this)
  }

  getContent () {
    const {
      type,
      properties,
      setProjectProperty,
      setExpeditionProperty,
      closeAndCancel,
      closeAndSave
    } = this.props
    switch (type) {
      case ('new project') : {
        return (
          <ModalNewProjectContent
            properties={ properties }
            setProjectProperty={ setProjectProperty }
            closeAndCancel={ closeAndCancel }
            closeAndSave={ closeAndSave }
          />
        )
      }
      case ('new expedition') : {
        return (
          <ModalNewExpeditionContent
            properties={ properties }
            setExpeditionProperty={ setExpeditionProperty }
            closeAndCancel={ closeAndCancel }
            closeAndSave={ closeAndSave }
          />
        )
      }
      default : {
        return null
      }
    }
  }

  render () {
    return (
      <div className="modal">
        { this.getContent() }
      </div>
    )    
  }
}

export default Modal