import '@mantine/core/styles.css';

import React from 'react'
import {createRoot} from 'react-dom/client'
import { MantineProvider } from '@mantine/core';
import { theme } from './theme';
import App from './app/App'

const container = document.getElementById('root')

const root = createRoot(container!)

root.render(
    <MantineProvider theme={theme} defaultColorScheme='dark'>
        <App/>
    </MantineProvider>
)
