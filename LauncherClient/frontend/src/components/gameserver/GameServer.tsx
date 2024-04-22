import React from 'react';
import s from './GameServer.module.css'
import Server from "./Server";
import {Server_Info} from '../../information/Server_Info'

const GameServer = () => {
    return (
        <div className={s.wrapper}>
            <h1 className={s.h1}>Список серверов</h1>
            <div className={s.servers}>

                {
                 Server_Info.map((obj) => (
                     <Server {...obj}/>
                 ))
                }

            </div>
        </div>
    );
};

export default GameServer;