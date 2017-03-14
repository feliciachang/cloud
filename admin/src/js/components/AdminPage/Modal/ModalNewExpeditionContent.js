
import React from 'react'
import slug from 'slug'

class ModalNewExpeditionContent extends React.Component {
  render () {
    const {
      projectName,
      properties,
      closeAndCancel,
      closeAndSave,
      setExpeditionProperty
    } = this.props

    const name = properties.get('name')
    const description = properties.get('description')

    return (
      <div>
        <div className="header">
          <h2>
            Create new expedition
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
                (!!name && name.toLowerCase() !== 'expedition name' ? '' : ' default')
              }
              value={name}
              onFocus={(e) => {
                if (!name || name === 'Expedition Name') {
                  setExpeditionProperty(['name'], '')
                }
              }}
              onChange={(e) => {
                setExpeditionProperty(['name'], e.target.value)
              }}
            />
            <p className="error"></p>
          </div>
        <div className="field">
          <p className="label">
            URL:
          </p>
          <p>
            { `https://${projectName}.fieldkit.org/${ slug(name) }` }
          </p>
        </div>
        <div className="field">
            <p className="label">
              Description:
            </p>
            <input 
              type="text"
              className={
                '' +
                (!!description && description.toLowerCase() !== 'expedition description' ? '' : ' default')
              }
              value={description}
              onFocus={(e) => {
                if (!description || description === 'Expedition Description') {
                  setExpeditionProperty(['description'], '')
                }
              }}
              onChange={(e) => {
                setExpeditionProperty(['description'], e.target.value)
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

export default ModalNewExpeditionContent