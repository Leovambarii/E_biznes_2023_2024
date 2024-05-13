const express = require('express');
const bodyParser = require('body-parser');
const Sequelize = require('sequelize');
const bcrypt = require('bcrypt');
const jwt = require('jsonwebtoken');
const cors = require('cors');
require('dotenv').config();

const app = express();
app.use(cors());
app.use(bodyParser.json());

const sequelize = new Sequelize({
  dialect: 'sqlite',
  storage: './Users.db',
  define: {
    timestamps: false,
  }
})

const User = sequelize.define('user', {
  username: {
    type: Sequelize.STRING,
    unique: true,
    allowNull: false
  },
  password: {
    type: Sequelize.STRING,
    allowNull: false
  }
});

sequelize.sync()
  .then(() => {
    console.log('Database synchronized');
  })
  .catch(err => {
    console.error('Database synchronization error:', err);
  });

  app.post('/register', async (req, res) => {
    const { username, password } = req.body;
    if (!username || !password) {
      return res.status(400).send('Username and password are required!');
    }

    try {
      const existingUser = await User.findOne({ where: { username } });
      if (existingUser) {
        return res.status(400).send('Username already taken!');
      }

      const hashedPassword = await bcrypt.hash(password, 10);
      const user = await User.create({username: username, password: hashedPassword });
      console.log('User registered:', user.username);
      res.status(201).send('Registration successful!');
    } catch (err) {
      console.error('User registration error:', err);
      res.status(500).send('Registration failed!');
    }
  });

  app.post('/login', async (req, res) => {
    const { username, password } = req.body;
    if (!username || !password) {
      return res.status(400).send('Username and password are required!');
    }

    try {
      const user = await User.findOne({ where: { username } });
      if (!user || !(await bcrypt.compare(password, user.password))) {
        res.status(401).send('Invalid login credentials!');
      } else {
        const token = jwt.sign({ userId: user.id, username: user.username }, process.env.SECRET_KEY);
        console.log('User logged in:', user.username);
        res.status(200).json({ token });
      }
    } catch (err) {
      console.error('User login error:', err);
      res.status(500).send('Login error!');
    }
  });

const port = 8080;

app.listen(port, () => {
  console.log('Server listening on port:', port);
});
