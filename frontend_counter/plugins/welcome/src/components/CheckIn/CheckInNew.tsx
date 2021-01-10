import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
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
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; 
import { EntCounterStaff, EntCustomer, EntDataRoom, EntReserveRoom } from '../../api';
import { Cookies } from '../../Cookie'
import { Link as RouterLink } from 'react-router-dom';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

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
  //CheckinDate : string;
  // create_by: number;
}

const CheckIn: FC<{}> = () => {
    
    const classes = useStyles();
    const api = new DefaultApi();

    // ดึงคุกกี้
    var ck = new Cookies()
    var cookieName = ck.GetCookie()

    //อันหลักสำหรับสร้าง การ Check_in
    const [CheckIn, setCheckIn] = React.useState<Partial<CheckIn>>({})
  
    // Customer
    const [idcustomer,setIdcustomer] = React.useState<number>(0)
    const [customer,setCustomer] = React.useState<EntCustomer[]>([])
    const getCustomer = async () => {
        const res = await api.listCustomer({})
        setCustomer(res)
    }
 
    // ReserveRoom  
    const [idreserveroom,setIdreserveroom] = React.useState<number>(0)
    const [reserveroom,setReserveroom] = React.useState<EntReserveRoom[]>([])
    const getReseveroom = async () => {
        const res = await api.getReserveRoomCustomer({id : idcustomer})
        setReserveroom(res)    
    }

    // DataRoom 
    const [dataroom,setDataroom] = React.useState<EntDataRoom[]>([])
    const getDataroom = async () => {
        const res = await api.getDataroomcustomer({id : idreserveroom})
        setDataroom(res)    
    }

  
    // CounterStaff
    const [counter,setCounter] = React.useState<EntCounterStaff[]>([])
    const getCounterStaff = async () => {
        const res = await api.listCounterStaff({})
        setCounter(res)
    }

    // handleChange
    const handleChange = (event: React.ChangeEvent<{ name?: string; value: any }>) => {
      const name = event.target.name as keyof typeof CheckIn
      const { value } = event.target
      setCheckIn({ ...CheckIn, [name]: value }) 
      setIdcustomer(event.target.value)
      setIdreserveroom(event.target.value)
    }

    const Toast = Swal.mixin({
        toast: true,
        position: 'top-end',
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        didOpen: toast => {
        toast.addEventListener('mouseenter', Swal.stopTimer);
        toast.addEventListener('mouseleave', Swal.resumeTimer);
        },
    });
  
    // Lifecycle Hooks
    useEffect(() => {
        getCustomer()
        getCounterStaff()
    }, [])
    useEffect(() => {
        getReseveroom()
    },[CheckIn.Customer]) 
    useEffect(() => {
        getDataroom()
    }, [CheckIn.Reserveroom]) 
    

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
        if (data.id != null) {
          //clear();
          Toast.fire({
            icon: 'success',
            title: 'Ckeck in สำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'error',
          });
        }
      });
  }

  function Clears() {
    ck.ClearCookie()
    window.location.reload(false)
  }
 
  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบย่อย Check In`}>
        <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10, marginRight:20 }}>{cookieName}</div>
        <Button
          variant="outlined"
          color="secondary"
          size="large"
          component={RouterLink}
          to="/"
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
                <InputLabel>Data Room</InputLabel>
                <Select
                  name="Dataroom"
                  value={CheckIn.Dataroom || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {dataroom.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.roomnumber}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            
            <Grid item xs={3}>
              <div className={classes.paper}>Counter staff</div>
            </Grid>
            <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Counter staff</InputLabel>
                <Select
                  name="Counter"
                  value={CheckIn.Counter || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {counter.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
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
        </Container>
      </Content>
    </Page>
  );
};

export default CheckIn;
