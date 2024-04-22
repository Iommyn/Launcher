import React, {useEffect, useState} from 'react';
import s from "./Home.module.css";
import {WindowCenter, WindowSetSize} from "../../../wailsjs/runtime";
import Sidebar from "../../components/sidebar/Sidebar";
import Slider from "../../components/slider/Slider";
import GameServer from "../../components/gameserver/GameServer";
import AccountDetails from "../../components/accountdetails/AccountDetails";
import Preloader from "../../components/UI/preloader/Preloader";


const Home = () => {

    const [isLoading, setIsLoading] = useState(true);

    const handleResize = () => {
        WindowSetSize(900, 500);
        WindowCenter();
    };

    useEffect(() => {
        handleResize();

        window.addEventListener('resize', handleResize);

        return () => {
            window.removeEventListener('resize', handleResize);
        };
    }, []);


    useEffect(() => {
        setTimeout(() => {
            setIsLoading(false);
        }, 1000);
    }, []);

    if (isLoading) {
        return <Preloader />;
    }


    return (
        <div className={s.container}>
            <Sidebar/>
            <div className={s.columns}>
                <div className={s.slider}>
                    <Slider/>
                </div>
                <div className={s.gameServer}>
                    <GameServer/>
                </div>
                <div className={s.AccountDetails}>
                    <AccountDetails/>
                </div>
            </div>

        </div>
    );
};

export default Home;