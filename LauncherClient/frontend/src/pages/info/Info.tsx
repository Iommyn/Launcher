import Sidebar from "../../components/sidebar/Sidebar";
import s from "./Info.module.css";
import Version from "../../components/versioninfo/Version";
import {useEffect, useState} from "react";
import Preloader from "../../components/UI/preloader/Preloader";
import ServerInfo from "../../components/serverinfopanel/ServerInfo";

const Info = () => {

    const [isLoading, setIsLoading] = useState<boolean>(true);

    useEffect(() => {
        const timer = setTimeout(() => {
            setIsLoading(false);
        }, 1000);

        return () => clearTimeout(timer);
    }, []);

    if (isLoading) {
        return <Preloader />;
    }

    return (
        <div className={s.container}>
            <Sidebar />
            <div className={s.columns}>
                <Version/>
                <ServerInfo/>
            </div>
        </div>
    );
};

export default Info;