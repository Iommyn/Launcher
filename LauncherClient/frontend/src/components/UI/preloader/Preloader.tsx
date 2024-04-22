import React from 'react';
import s from './Preloader.module.css'

const Preloader = () => {
    return (
        <div className={s.preloader}>
            <div className={s.loader}>
            </div>
            <h1 className={s.h1}>Загрузка....</h1>
        </div>
    );
};

export default Preloader;