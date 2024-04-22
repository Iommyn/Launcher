import React, { useState, useEffect } from 'react';
import s from './Slider.module.css';
import {slides} from '../../information/Slider';

const Slider = () => {


    const [currentSlide, setCurrentSlide] = useState(0);

    useEffect(() => {
        const interval = setInterval(() => {
            setCurrentSlide((prevSlide) => (prevSlide === slides.length - 1 ? 0 : prevSlide + 1));
        }, 5000);

        return () => clearInterval(interval);
    }, [slides.length]);

    return (
        <>
            <div className={s.slider}>
                <div className={s.slideContent}>
                    <div className={s.slideText}>
                        <h1 className={s.title}>{slides[currentSlide].title}</h1>
                        <h2 className={s.description} dangerouslySetInnerHTML={{ __html: slides[currentSlide].description }}></h2>
                    </div>
                    <div className={s.slideImage}>
                        <img src={slides[currentSlide].image} alt={slides[currentSlide].title} />
                    </div>
                </div>
            </div>
            <div className={s.pagination}>
                {slides.map((_, index) => (
                    <span key={index} className={index === currentSlide ? s.dot + ' ' + s.activeDot : s.dot}></span>
                ))}
            </div>
        </>
    );
};

export default Slider;