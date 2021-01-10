import '@backstage/cli/asset-types';
import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import Log from './Log';
import { Cookies } from './Cookie'

var ck = new Cookies()
var cookie = ck.GetCookie()

if(cookie == undefined){
    ReactDOM.render(<Log />, document.getElementById('root'));
}
else{
    ReactDOM.render(<App />, document.getElementById('root'));
}

