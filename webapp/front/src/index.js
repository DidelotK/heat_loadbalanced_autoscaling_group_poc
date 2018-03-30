import React from 'react';
import ReactDOM from 'react-dom';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Dashboard from './Dashboard';


const App = () => (
  <MuiThemeProvider>
    <Dashboard />
  </MuiThemeProvider>
);

ReactDOM.render(<App />, document.getElementById('root'));
