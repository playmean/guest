/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./index.html', './src/**/*.{vue,ts,js}'],
    theme: {
        extend: {
            colors: {
                'guest-dark': 'hsl(0, 0%, 7%)',
                'guest-lighten': 'hsl(240, 2%, 11%)',
                'guest-light': 'hsl(240, 3%, 18%)',
                'guest-slate': 'hsl(230, 6%, 20%)',
                'guest-accent': 'hsl(253, 52%, 58%)',
            },
        },
    },
    plugins: [],
};
