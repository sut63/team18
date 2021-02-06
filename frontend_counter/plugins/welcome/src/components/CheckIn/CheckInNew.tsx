import React, { FC, useEffect } from 'react';
import { Theme, makeStyles, withStyles, createStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import Swal from 'sweetalert2'; // alert

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  Avatar,
  Button,
  TextField,
  Badge,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
import { EntCounterStaff, EntCustomer, EntDataRoom, EntReserveRoom } from '../../api';
import { Cookies } from '../../Cookie'
import { Link as RouterLink } from 'react-router-dom';
import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';
import accountImg from '../../image/account.jpg'

// header css
const HeaderCustom = {
  minHeight: '50px',
};

const StyledBadge = withStyles((theme: Theme) =>
  createStyles({
    badge: {
      backgroundColor: '#44b700',
      color: '#44b700',
      boxShadow: `0 0 0 2px ${theme.palette.background.paper}`,
      '&::after': {
        position: 'absolute',
        top: 0,
        left: 0,
        width: '100%',
        height: '100%',
        borderRadius: '50%',
        animation: '$ripple 1.2s infinite ease-in-out',
        border: '1px solid currentColor',
        content: '""',
      },
    },
    '@keyframes ripple': {
      '0%': {
        transform: 'scale(.8)',
        opacity: 1,
      },
      '100%': {
        transform: 'scale(2.4)',
        opacity: 0,
      },
    },
  }),
)(Badge);

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

interface CheckIn {
  Customer: number;
  Counter: number;
  Reserveroom: number;
  Dataroom: number;
  MobileKey: string;
  PhoneNumber: string;
  PersonNumber: string;
  //CheckinDate : string;
  // create_by: number;
}

const CheckIn: FC<{}> = () => {

  const classes = useStyles();
  const api = new DefaultApi();

  // ดึงคุกกี้
  var ck = new Cookies()
  var cookieName = ck.GetCookie()
  var cookieID = ck.GetID()

  //อันหลักสำหรับสร้าง การ Check_in
  const [CheckIn, setCheckIn] = React.useState<Partial<CheckIn>>({})

  // สำหรับตรวยสอบความถูกต้อง
  const [MobileKeyError, setMobileKeyError] = React.useState('');
  const [PhoneNumberError, setPhoneNumberError] = React.useState('');
  const [PersonNumberError, setPersonNumberError] = React.useState('');
  const [errors, setError] = React.useState(String);

  // Customer
  const [idcustomer, setIdcustomer] = React.useState<number>(0)
  const [customer, setCustomer] = React.useState<EntCustomer[]>([])
  const getCustomer = async () => {
    const res = await api.listCustomer({})
    setCustomer(res)
  }

  // ReserveRoom  
  const [idreserveroom, setIdreserveroom] = React.useState<number>(0)
  const [reserveroom, setReserveroom] = React.useState<EntReserveRoom[]>([])
  const getReseveroom = async () => {
    const res = await api.getReserveRoomCustomer({ id: idcustomer })
    setReserveroom(res)
  }

  // DataRoom 
  const [dataroom, setDataroom] = React.useState<EntDataRoom>()
  const getDataroom = async () => {
    const res = await api.getDataroomcustomer({ id: idreserveroom })
    setDataroom(res)
  }

  // CounterStaff
  const [counter, setCounter] = React.useState<EntCounterStaff>()
  const getCounterStaff = async () => {
    const res = await api.getCounterStaff({ id: Number(cookieID) })
    setCounter(res)
  }

  // handleChange
  const handleChange = (event: React.ChangeEvent<{ name?: string; value: any }>) => {
    const name = event.target.name as keyof typeof CheckIn
    const { value } = event.target
    setCheckIn({ ...CheckIn, [name]: value })
    setIdcustomer(event.target.value)
    setIdreserveroom(event.target.value)
    const validateValue = value.toString()
    checkPattern(name, validateValue)
  }

  //valid
  const validateMobileKey = (val: string) => {
    return val.length == 10 ? true : false;
  }
  const validatePhoneNumber = (val: string) => {
    return val.match("[0]\\d{9}") && val.length == 10 ? true : false;
  }
  const validatePersonNumber = (val: string) => {
    return val.match("[0-9]\\d{12}") && val.length == 13 ? true : false;
  }

  // checkPattern
  const checkPattern = (id: string, value: string) => {
    switch (id) {
      case 'MobileKey':
        validateMobileKey(value) ? setMobileKeyError('') : setMobileKeyError('ความปลอดภัยต่ำ กรุณาใส่ 10 ตัว');
        return;
      case 'PhoneNumber':
        validatePhoneNumber(value) ? setPhoneNumberError('') : setPhoneNumberError('Ex 0600000000');
        return;
      case 'PersonNumber':
        validatePersonNumber(value) ? setPersonNumberError('') : setPersonNumberError('Ex 1110000000000')
        return;
      default:
        return;
    }
  }

  // alert setting
  const [open, setOpen] = React.useState(false);
  const [fail, setFail] = React.useState(false);

  //close alert 
  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }
    setFail(false);
    setOpen(false);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getCustomer()
    getCounterStaff()
  }, [])
  useEffect(() => {
    getReseveroom()
  }, [CheckIn.Customer])
  useEffect(() => {
    getDataroom()
  }, [CheckIn.Reserveroom])
  useEffect(() => {
    setCheckIn({ ...CheckIn, ['Counter']: counter?.id })
  }, [counter]);
  useEffect(() => {
    setCheckIn({ ...CheckIn, ['Dataroom']: dataroom?.id })
  }, [dataroom]);

  // func checkerror
  const checkerror = (s: string) => {
    switch (s) {
      case 'phone_number':
        setError("เบอร์โทรต้องเป็นตัวเลข 10 หลักและขึ้นต้นด้วย 0")
        return;
      case 'mobile_key':
        setError("Mobile Key ต้องมีความยาว 10 ตัว")
        return;
      case 'person_number':
        setError("เลขบัตรประชาชนต้องเป็นตัวเลข 13 หลัก")
        return;
      default:
        setError("กรุณากรอกข้อมูลให้ครบถ้วน")
        return;
    }
  };



  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/checkins';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(CheckIn),
    };

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          clear();
          setOpen(true);
        } else {
          checkerror(data.error.Name)
          setFail(true);
        }
      });
  }

  function Clears() {
    ck.ClearCookie()
    window.location.reload(false)
  }

  // clear input form
  function clear() {
    setCheckIn({});
    getCounterStaff();
  }

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบย่อย Check In`}>
        <StyledBadge
          overlap="circle"
          anchorOrigin={{
            vertical: 'bottom',
            horizontal: 'right',
          }}
          variant="dot"
        >
          <Avatar alt="Remy Sharp" src={accountImg} />
        </StyledBadge>
        <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
        <Button
          variant="outlined"
          color="secondary"
          size="large"
          onClick={Clears}
        >
          Logout
        </Button>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>customer</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>select customer</InputLabel>
                <Select
                  name="Customer"
                  value={CheckIn.Customer || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {customer.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>reserveroom</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>select reserveroom</InputLabel>
                <Select
                  name="Reserveroom"
                  value={CheckIn.Reserveroom || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {reserveroom.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.id}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>Room Number</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  disabled
                  name="Dataroom"
                  variant="outlined"
                  value={dataroom?.roomnumber}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>Counter staff</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  disabled
                  name="Counter"
                  variant="outlined"
                  value={counter?.name}
                />
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>Peson id</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  error={PersonNumberError ? true : false}
                  helperText={PersonNumberError}
                  name="PersonNumber"
                  label="เลขบัตรประชาชน"
                  variant="outlined"
                  value={CheckIn.PersonNumber || ''}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>Phone number</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  error={PhoneNumberError ? true : false}
                  helperText={PhoneNumberError}
                  name="PhoneNumber"
                  label="เบอร์โทรศัพท์"
                  variant="outlined"
                  value={CheckIn.PhoneNumber || ''}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>Mobile Key</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  error={MobileKeyError ? true : false}
                  helperText={MobileKeyError}
                  name="MobileKey"
                  label="คีย์เข้าห้อง"
                  variant="outlined"
                  value={CheckIn.MobileKey || ''}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>
            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                onClick={save}
              >
                Check in
              </Button>
            </Grid>
          </Grid>

          <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="success">
              บันทึกสำเร็จ
          </Alert>
          </Snackbar>

          <Snackbar open={fail} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="error">
              {errors}
            </Alert>
          </Snackbar>

        </Container>
      </Content>
    </Page>
  );
};

export default CheckIn;
