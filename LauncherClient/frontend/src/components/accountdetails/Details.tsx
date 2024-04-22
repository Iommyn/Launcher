import React from 'react';
import s from "./AccountDetails.module.css";
import {AccountItem} from "../../types/types";

const Details = ({img, title, content}:AccountItem) => {
    return (
        <div className={s.details}>
            <div className={s.wrapper_content}>
            <div className={s.container}>
                <div>
                    <img src={img}/>
                </div>
                <h1 className={s.title}>{title}</h1>
            </div>
            <h1 className={s.title_content}>{content}</h1>
            </div>
        </div>
    );
};

export default Details;