import React, { Component } from 'react';
import axios from 'axios';
import Paper from 'material-ui/Paper';
import RaisedButton from 'material-ui/RaisedButton';
import TextField from 'material-ui/TextField';
import './dashboard.css';
import {
  PieChart,
  Pie,
  Tooltip,
  ResponsiveContainer, Cell,
} from 'recharts';
axios.defaults.baseURL = 'http://localhost:8080/api';


const COLORS = ['#FFF', '#FF4081'];

const styles = {
  ipContainer: {
    height: 210,
    width: 350,
    margin: 50,
    textAlign: 'center',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    flexDirection: 'column',
    fontSize: 30
  },
  makeLoadContainer: {
    height: 210,
    width: 350,
    margin: 50,
    textAlign: 'center',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    flexDirection: 'column'
  }
};


class Dashboard extends Component {
  constructor(props) {
    super(props);
    this.state = {
      cpu: null,
      ip: '',
      makeCpuLoadDuration: 10
    };

    this._getCpu = this._getCpu.bind(this);
    this._getIp = this._getIp.bind(this);
    this._makeCpuLoad = this._makeCpuLoad.bind(this);
    this._handleChangeDuration = this._handleChangeDuration.bind(this);
  }

  componentWillMount() {
    this._getCpu();
    this._getIp();
  }

  componentDidMount() {
    this.timer = setInterval(() => this._getCpu(), 1000);
  }

  componentWillUnmount() {
    clearInterval(this.timer);
  }

  _getCpu() {
    axios.get('/cpu')
      .then(res => this.setState({ cpu: Number(res.data.toFixed(1)) }));
  }

  _getIp() {
    axios.get('/ip')
      .then(res => this.setState({ ip: res.data }));
  }

  _makeCpuLoad() {
    axios.get(`/makeCpuLoad?duration=${this.state.makeCpuLoadDuration}`);
  }

  _handleChangeDuration(e, value) {
    if (value > 0 && value <= 60) {
      this.setState({ makeCpuLoadDuration: value });
    }
  }

  render() {
    const pieData = [
      {
        name: 'Cpu',
        value: 100 - this.state.cpu
      },
      {
        name: 'Empty',
        value: this.state.cpu
      }
    ];

    return (
      <div id="Dashboard">
        <h1>Poc HEAT IaaC</h1>

        <section className="topZone">
          <Paper style={styles.ipContainer} zDepth={1}>
            <div><b>IP</b></div>
            <div>{this.state.ip}</div>
          </Paper>

          <Paper style={styles.makeLoadContainer} zDepth={1}>
            <div className="makeLoadLabel"><b>Make cpu load duration</b></div>
            <TextField
              type="number"
              hintText="Duration (in sec)"
              value={this.state.makeCpuLoadDuration}
              onChange={this._handleChangeDuration}
            /><br />
            <RaisedButton label="Make cpu load" secondary={true} onClick={this._makeCpuLoad}/>
          </Paper>
        </section>

        {this.state.cpu !== null ?
          <div className="cpuIndicator">
            <div className="cpuLabel">
              <b>Cpu:</b> {this.state.cpu}%
            </div>
            <ResponsiveContainer>
              <PieChart>
                <Pie
                  data={pieData}
                  nameKey="name"
                  dataKey="value"
                  fill="#82ca9d"
                  labelLine={false}
                >
                  {
                    pieData.map((entry, index) => <Cell key={index} fill={COLORS[index]}/>)
                  }
                </Pie>
                <Tooltip />
              </PieChart>
            </ResponsiveContainer>
          </div>
          : null
        }
      </div>
    );
  }
}

export default Dashboard;
