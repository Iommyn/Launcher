import React, { useState } from 'react';
import s from "./Version.module.css";


type GameVersion = '1.12.2' | '1.7.2';

const Version = () => {
    const [activeVersion, setActiveVersion] = useState<GameVersion>('1.12.2');

    const handleVersionClick = (version: GameVersion) => {
        setActiveVersion(version);
    };

    return (
        <div className={s.version}>
            <h1 className={s.title_version}>Версия игры</h1>
            <button
                className={`${s.btn_version}  ${activeVersion === '1.12.2' ? s.active : ''}`}
                onClick={() => handleVersionClick('1.12.2')}
            >
                1.12.2
            </button>
            <button
                className={`${s.btn_versionTwo} ${activeVersion === '1.7.2' ? s.active : ''}`}
                onClick={() => handleVersionClick('1.7.2')}
            >
                1.7.2
            </button>
        </div>
    );
};

export default Version;