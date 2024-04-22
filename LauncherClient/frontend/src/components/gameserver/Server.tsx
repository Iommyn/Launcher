import React from 'react';
import s from "./GameServer.module.css";
import {Servers} from "../../types/types";

const Server = ({title, online}:Servers) => {
    return (
            <div className={s.server}>
                <h1>{title}</h1>
                <div className={s.infoServer}>
                    <h2>{online}</h2>
                    <div className={s.progress}></div>
                </div>
            </div>
    );
};

export default Server;