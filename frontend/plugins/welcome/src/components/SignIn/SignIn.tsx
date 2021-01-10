import React, { FC , useEffect } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { Cookies } from '../../Cookie';
import { DefaultApi } from '../../api/apis'; 
import { EntCounterStaff, EntCustomer } from '../../api';
import { ApiProvider } from '@backstage/core';
import { Link as RouterLink } from 'react-router-dom';

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Your Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const SignIn: FC<{}> = () => {

  const classes = useStyles();
  const api = new DefaultApi();
  var ck = new Cookies();
  var check : boolean
  const [path, setPath] = React.useState("");

  // list CounterStaff
  const [counter,setCounter] = React.useState<EntCustomer[]>([])
  const listCounter = async() => {
        const res = await api.listCustomer({})
        setCounter(res)
  }

  // setEmail
  const [email, setEmail] = React.useState()
  const handleEmail = (event : any) => {
      setEmail(event.target.value)
  }

  // setPassword
  const [password, setPassword] = React.useState()
  const handlePassword = (event : any) => {
      setPassword(event.target.value)
  }

  // handleCookies
  function handleCookies() {
    check = ck.CheckLogin(counter,email,password)
    console.log("check => "+check)
    if(check === true){
      setPath("/WelcomePage")
      ck.SetCookie("user_email",email,30)
      ck.SetCookie("user_id",ck.setID(counter,email,password),30)
      window.location.reload(false)
    }else if(check === false){
      alert("The wrong password or email was entered.!!!")
      setPath("/")
    }
  }
  // useEffect
  useEffect(() => {
      listCounter()
  },[])

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <form className={classes.form} noValidate>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
            onChange={handleEmail}
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            onChange={handlePassword}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={handleCookies}
            component={RouterLink}
            to={path}
          >
            Sign In
          </Button>
        </form>
      </div>
    </Container>
  );
};

export default SignIn;
