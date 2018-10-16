import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  state = {
    tasks: []
  }

  getTasks = async () => {
    const res = await fetch('/tasks')
    const json = await res.json()
    return json.items
  }

  async componentDidMount() {
    // TODO: replace this with context or redux or whatever
    const tasks = await this.getTasks()
    this.setState({ tasks })
  }

  render() {
    const { tasks } = this.state

    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.js</code> and save to reload.
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>

          {tasks.map(task => (
            <div key={task.id}>
              <strong>{task.name}</strong>
            </div>
          ))}
        </header>
      </div>
    );
  }
}

export default App;
