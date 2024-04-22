import React, { useState, useEffect, FormEvent } from 'react';
import { useNavigate } from 'react-router-dom';
import { BrowserOpenURL, WindowSetSize, WindowCenter } from "../../../wailsjs/runtime";
import s from './Auth.module.css';
import Input from '../../components/UI/input/Input'

const getCookie = (name: string): string | undefined => {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop()?.split(';').shift();
};

const Auth: React.FC = (props) => {
    const navigate = useNavigate();
    const [nick, setNick] = useState<string>(getCookie('nick') || '');
    const [password, setPassword] = useState<string>(getCookie('password') || '');
    const [remember, setRemember] = useState<boolean>(false);


    function handleResetPass() {
        BrowserOpenURL('https://corecraft.ru/') //Ссылка на восстановление пароля corecraft/reset
    }

    function handleReg() {
        BrowserOpenURL('https://corecraft.ru/') //Ссылка на регистрацию corecraft/register
    }

    function handleHelp() {
        BrowserOpenURL('https://corecraft.ru/') //Ссылка на форум
    }



    const setCookie = (name: string, value: string, days: number): void => {
        let expires = "";
        if (days) {
            const date = new Date();
            date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
            expires = "; expires=" + date.toUTCString();
        }
        document.cookie = name + "=" + (value || "") + expires + "; path=/";
    };

    const deleteCookie = (name: string): void => {
        document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
    };


    const handleSubmit = (e: FormEvent<HTMLFormElement>): void => {
        e.preventDefault();
        if (remember) {
            setCookie('nick', nick, 30);
            setCookie('password', password, 30);
        } else {
            deleteCookie('nick');
            deleteCookie('password');
        }
        navigate('/home');
    };

    const handleRememberChange = (): void => {
        setRemember(prevRemember => !prevRemember);
    };


    const handleResize = (): void => {
        WindowSetSize(522, 568);
        WindowCenter();
    };


    useEffect(() => {
        handleResize();
        window.addEventListener('resize', handleResize);
        return () => {
            window.removeEventListener('resize', handleResize);
        };
    }, []);

    return (
        <div className={s.section}>
            <div className={s.wrapper}>
                <div>
                    <h1 className={s.h1}>Авторизация</h1>
                    <form onSubmit={handleSubmit}>
                        <div className={s.container}>
                            <Input
                                type="text"
                                placeholder="Введите никнейм"
                                value={nick}
                                onChange={(e) => setNick(e.target.value)}
                            />
                            <Input
                                type="password"
                                placeholder="Введите Пароль"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                            />
                            <div className={s.checkbox}>
                                <input type="checkbox" id="check" onChange={handleRememberChange} checked={remember}/>
                                <label htmlFor="check">Запомнить меня</label>
                            </div>
                            <button className={s.btn} type="submit">
                                Войти
                            </button>

                            <h2 onClick={handleReg} className={s.h2}>Нет аккаунта? <span
                                className={s.span}>Зарегистрируйтесь</span></h2>
                            <h2 onClick={handleResetPass} className={s.recovery}>Восстановить аккаунт</h2>
                            <h2 onClick={handleHelp} className={s.problems}>Возникла проблема?</h2>


                        </div>
                    </form>
                </div>
            </div>
        </div>
    );
};

export default Auth;