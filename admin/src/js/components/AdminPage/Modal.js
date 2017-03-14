
import React from 'react'

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
      closeModalAndCancel,
      closeModalAndSave
    } = this.props
    switch (type) {
      case ('new project') : {
        const name = properties.get('name')
        const description = properties.get('description')

        return (
          <div>
            <div className="header">
              <h2>
                Create new project
              </h2>
            </div>
            <div className="content">
              <div className="field">
                <p className="label">
                  Name:
                </p>
                <input 
                  type="text"
                  className={
                    '' +
                    (!!name && name.toLowerCase() !== 'project name' ? '' : ' default')
                  }
                  value={name}
                  onFocus={(e) => {
                    if (!name || name === 'Project Name') {
                      setProjectProperty(['name'], '')
                    }
                  }}
                  onChange={(e) => {
                    setProjectProperty(['name'], e.target.value)
                  }}
                />
                <p className="error"></p>
              </div>
              <div className="field">
                <p className="label">
                  Description:
                </p>
                <input 
                  type="text"
                  className={
                    '' +
                    (!!description && description.toLowerCase() !== 'project description' ? '' : ' default')
                  }
                  value={description}
                  onFocus={(e) => {
                    if (!description || description === 'Project Description') {
                      setProjectProperty(['description'], '')
                    }
                  }}
                  onChange={(e) => {
                    setProjectProperty(['description'], e.target.value)
                  }}
                />
                <p className="error"></p>
              </div>
            </div>
            <div className="actions">
              <div
                className="button"
                onClick={ closeModalAndCancel }
              >
                Cancel
              </div>
              <div
                className="button"
                onClick={ closeModalAndSave }
              >
                Save
              </div>
            </div>
          </div>
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