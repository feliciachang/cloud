
import React from 'react'

class ModalNewProjectContent extends React.Component {
  render () {

    const {
      properties,
      closeAndCancel,
      closeAndSave,
      setProjectProperty
    } = this.props

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
            onClick={ closeAndCancel }
          >
            Cancel
          </div>
          <div
            className="button"
            onClick={ closeAndSave }
          >
            Save
          </div>
        </div>
      </div>
    )
  }
}

export default ModalNewProjectContent