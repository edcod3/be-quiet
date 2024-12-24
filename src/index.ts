import express, { Request, Response } from 'express';
import path from 'path';

const app = express();
const port = 3000;

// Serve static files from the 'dist' folder
app.use(express.static(path.join(__dirname, 'public')));
app.use('/styles', express.static(path.join(__dirname, 'styles')));

// Basic route
app.get('/', (req: Request, res: Response) => {
  res.sendFile(path.join(__dirname, 'public', 'index.html'));
});

app.listen(port, () => {
  console.log(`Server is running at http://localhost:${port}`);
});

app.get('/api/alert', (req: Request, res: Response) => {
    console.log("Received alert:", req.params);
    res.json({'msg': 'Alert sent!'});
});