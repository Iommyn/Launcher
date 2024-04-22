
import s from './Input.module.css'
import { useState } from 'react';
import { InputHTMLAttributes, ForwardRefExoticComponent, forwardRef, Ref } from 'react';

interface MyInputProps extends InputHTMLAttributes<HTMLInputElement> {}

const Input: ForwardRefExoticComponent<MyInputProps> = forwardRef(
    (props: MyInputProps, ref: Ref<HTMLInputElement>) => {
        const [showPassword, setShowPassword] = useState(false);

        const togglePasswordVisibility = () => {
            setShowPassword(!showPassword);
        };

        return (
            <div style={{ position: 'relative' }}>
                <input
                    type={showPassword ? 'text' : 'password'}
                    ref={ref}
                    className={s.myInput}
                    style={{ outline: 'none' }}
                    {...props}
                />

            </div>
        );
    }
);

export default Input;