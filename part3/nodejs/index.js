const express = require('express');
const app = express();

const port = 4000;

app.get('/', (req, res) => {
    res.send('===== APP CON NODEJS Y DOCKERFILE =====');
})

app.listen(port, () => {
    console.log(`Servidor corriendo en el puerto ${port}`);
});