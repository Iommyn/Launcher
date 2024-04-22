import React from 'react';
import s from './AccountDetails.module.css';
import Details from "./Details";
import {Account} from "../../information/Account";

const AccountDetails = () => {
    return (
        <div className={s.wrapper}>
            <h1 className={s.h1}>Аккаунт</h1>
            {
                Account.map((obj) => (
                    <Details {...obj}/>
                ))
            }
            <button className={s.button}>Выйти из аккаунта</button>
        </div>
    );
};

export default AccountDetails;