import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.props = props;
    this.state = {
      lat: 49.2827,
      long: -123.1207,
      roads: []
    }

    this.onSubmit = this.onSubmit.bind(this);
    this.handleChange = this.handleChange.bind(this);
  }

  async onSubmit(e) {
    e.preventDefault();
    const response = await fetch(`http://localhost:8000/build_game?lat=${this.state.lat}&lon=${this.state.long}`);
    const res = await response.json();
    this.setState({ roads: res.rooads });
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
          <CanvasComponent roads={this.state.roads} />
        </div>
      </div>
    );
  }
}

class CanvasComponent extends React.Component {
  canvas;

  constructor(props) {
    super(props);
    this.props = props;
  }
  
  componentDidMount() {
      this.canvas = this.refs.canvas.getContext('2d');
      this.initCanvas();
  }

  initCanvas() {
    this.canvas.fillRect(0, 0, 400, 400);
    this.canvas.strokeStyle = 'orange';
  }

  updateCanvas(ctx, roads) {
    if (!ctx || !roads) return;
    console.log('updating...');

    roads.forEach((road) => {
      const coordinates = [{ latitude: 50, longitude: 200 }, { latitude: 250, longitude: 200 }, { latitude: 250, longitude: 10 }, { latitude: 350, longitude: 10 }] || road.coordinates;
      coordinates.forEach((c) => {
        ctx.lineTo(c.latitude, c.longitude);
        ctx.moveTo(c.latitude, c.longitude);
        ctx.stroke();
      });
    });
  }

  render() {
    this.updateCanvas(this.canvas, this.props.roads);
    return (
      <canvas className="Board" ref="canvas" width={400} height={400} />
    );
  }
}

export default App;
