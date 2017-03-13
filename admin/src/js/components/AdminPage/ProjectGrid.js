
import React from 'react'
import { Link } from 'react-router'

class ProjectGrid extends React.Component {
  
  render () {

    const {
      projects
    } = this.props

    return (
      <div id="project-grid" className="project-grid">
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
    )
  }
}

export default ProjectGrid