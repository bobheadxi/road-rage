import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.props = props;
    this.state = {
      lat: 49.2827,
      long: 123.1207
    }

    this.onSubmit = this.onSubmit.bind(this);
    this.handleChange = this.handleChange.bind(this);
  }

  onSubmit(e) {
    e.preventDefault();
    console.log(this.state);
  }

  handleChange(e, l) {
    if (l === 'lat') {
      this.setState({ lat: e.target.value });
    } else if (l === 'long') {
      this.setState({ long: e.target.value });
    }
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to Road Rage</h1>
        </header>
        <p className="App-intro">
          To get started, enter a set of coordinates.
        </p>
        <form style={{ marginBottom: 20 }} onSubmit={this.onSubmit}>
          <input type="text" name="lat" placeholder="enter latitude" value={this.state.lat} onChange={(e) => this.handleChange(e, 'lat')} />
          <input type="text" name="long" placeholder="enter longitude" value={this.state.long} onChange={(e) => this.handleChange(e, 'long')} />
          <input type="submit" name="submit" onClick={this.handleChange} />
        </form>
        <div className="Board-wrapper" style={{ width: 400, height: 400 }}>
          <CanvasComponent />
        </div>
      </div>
    );
  }
}

class CanvasComponent extends React.Component {
  componentDidMount() {
      this.updateCanvas();
  }

  updateCanvas() {
      const ctx = this.refs.canvas.getContext('2d');
      ctx.fillRect(0, 0, 400, 400); // background hack lol
  }

  render() {
      return (
          <canvas className="Board" ref="canvas" width={400} height={400} />
      );
  }
}

export default App;
