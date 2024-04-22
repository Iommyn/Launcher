import React, { useState, useEffect } from 'react';
import { Link, useLocation } from 'react-router-dom';
import s from './Sidebar.module.css';
import home from '../../assets/images/home.png';
import homeActive from '../../assets/images/home_active.png';
import info from '../../assets/images/info.png';
import infoActive from '../../assets/images/info_active.png';


type ActiveItem = 'home' | 'info';

const Sidebar: React.FC = () => {
    const location = useLocation();
    const [activeItem, setActiveItem] = useState<ActiveItem>('home');

    useEffect(() => {
        if (location.pathname === '/home') {
            setActiveItem('home');
        } else if (location.pathname === '/info') {
            setActiveItem('info');
        }
    }, [location]);

    return (
        <div className={s.sidebar}>
            <div className={s.wrapper}>
                <div>
                    <Link to="/home" className={s.item}>
                        <img className={s.home} src={activeItem === 'home' ? homeActive : home} alt="home" />
                    </Link>
                </div>

                <div>
                    <Link to="/info" className={s.item}>
                        <img className={s.info} src={activeItem === 'info' ? infoActive : info} alt="info" />
                    </Link>
                </div>
            </div>
        </div>
    );
};

export default Sidebar;