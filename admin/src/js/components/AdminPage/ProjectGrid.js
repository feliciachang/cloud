
import React from 'react'
import { Link } from 'react-router'

class ProjectGrid extends React.Component {
  
  render () {

    const {
      projects,
      user,
      promptModalNewProject
    } = this.props

    console.log('aga', user.toJS())

    return (
      <div id="project-grid" className="project-grid">
        <div className="greetings">
          <img/>
          <h2>
            {`Welcome back, ${ user.get('username') }`}
          </h2>
        </div>
        <div className="projects">
          <p>
            Choose a project
          </p>
          <div className="grid">
            { 
              projects.map((p, i) => {
                return (
                  <Link
                    to={ '/admin/' + p.get('id') }
                  >
                    <h2
                      key={ i }
                    >
                      { p.get('name') }
                    </h2>
                  </Link>
                )
              })
            }
          </div>
          <div
            className="button"
            onClick={ promptModalNewProject }
          >
            Create new project
          </div>
        </div>
      </div>
    )
  }
}

export default ProjectGrid