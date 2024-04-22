import React, { useState } from 'react';
import s from "./ServerInfo.module.css";
import { Link } from "react-router-dom";

interface ServerInfoType {
    title: string;
    description: string;
}

interface ServerData {
    [key: string]: ServerInfoType;
}

const ServerInfo = () => {
    const [activeServer, setActiveServer] = useState<string>('TechnoMagic');

    const serverInfo: ServerData = {
        'TechnoMagic': {
            title: 'TECHNOMAGIC',
            description: 'Свежий взгляд на техническую и магическую сборку с модами. Играть можно уже сейчас!',
        },
        'Hi-Tech': {
            title: 'Hi-Tech',
            description: 'укнукнукнHi-вера.Свежий взгляд на технино уже сейчас!',
        },
        'Magic': {
            title: 'Magic',
            description: 'Описание для Hi-вера.Свежий взгляд на технино уже сейчас!цукецукцу',
        },
    };

    const handleServerChange = (server: string) => {
        setActiveServer(server);
    };

    return (
        <div className={s.ServerInfo}>
            <div className={s.wrapper}>
                <h1 className={s.title}>{serverInfo[activeServer].title}</h1>
                <h2 className={s.description}>{serverInfo[activeServer].description}</h2>
                <div>
                    <button className={s.button}>Играть</button>
                    <Link to={'/'} className={s.link}>Настройки</Link>
                </div>
            </div>
            <div className={s.container}>
                {Object.keys(serverInfo).map((server, index) => (
                    <div
                        key={index}
                        className={`${s.server} ${activeServer === server ? s.active : ''}`}
                        onClick={() => handleServerChange(server)}
                    >
                        {server}
                    </div>
                ))}
            </div>
        </div>
    );
};

export default ServerInfo;