import '@backstage/cli/asset-types';
import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import Log from './Log';
import { Cookies } from 'plugin-welcome/src/Cookie'

var ck = new Cookies()
var role = ck.GetRole()

if(role != "customer"){
    ReactDOM.render(<Log />, document.getElementById('root'));
}
else{
    ReactDOM.render(<App />, document.getElementById('root'));
}

