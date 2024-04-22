import React, {useEffect} from 'react'
import {createRoot} from 'react-dom/client'
import logo from './assets/images/logo.png'
import App from './App'
import s from './style.module.css'
import {BrowserOpenURL, Quit} from "../wailsjs/runtime";

const container = document.getElementById('root')

const root = createRoot(container!)

const handleClose = () => {
    Quit()
}


function handleSiteLink() {
    BrowserOpenURL('https://corecraft.ru/')
}

root.render(
    <React.StrictMode>
        <div className={s.logo}>
            <img onClick={handleSiteLink} src={logo}/>
        </div>
        <div className={s.handleClose} onClick={handleClose}>X</div>
        <App/>
    </React.StrictMode>
)



