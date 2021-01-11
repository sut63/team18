import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import { Cookies } from '../../Cookie' 

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
import { EntPromotion } from '../../api/models/EntPromotion'; 
import { EntTypeRoom } from '../../api/models/EntTypeRoom';
import { EntStatusRoom } from '../../api/models/EntStatusRoom'; 

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

interface DataRoom {
  Promotion: number;
  StatusRoom: number;
  TypeRoom: number;
  Price: number;
  RoomNumber: string;
  // create_by: number;
}

const DataRoom: FC<{}> = () => {
  const classes = useStyles()
  const api = new DefaultApi()
  var ck = new Cookies()
  var cookieName = ck.GetCookie()
  const [DataRoom, setDataRoom] = React.useState<Partial<DataRoom>>({});
  const [StatusRoom, setStatusRoom] = React.useState<EntStatusRoom[]>([]);
  const [Promotion, setPromotion] = React.useState<EntPromotion[]>([]);
  const [TypeRoom, setTypeRoom] = React.useState<EntTypeRoom[]>([]);

  // alert setting
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

  const getStatusRoom = async () => {
    const res = await api.listStatusroom({ limit: 10, offset: 0 });
    setStatusRoom(res);
  };

  const getTypeRoom = async () => {
    const res = await api.listTyperoom({ limit: 10, offset: 0 });
    setTypeRoom(res);
  };
  const getPromotion = async () => {
    const res = await api.listPromotion({ limit: 10, offset: 0 });
    setPromotion(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getStatusRoom();
    getPromotion();
    getTypeRoom();
  }, []);

  // set data to object DataRoom
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof DataRoom;
    const { value } = event.target;
    setDataRoom({ ...DataRoom, [name]: value });
    console.log(DataRoom);
  };
  var Price  =  DataRoom.Price
  DataRoom.Price = Number(Price)
  
  
  // clear input form
  function clear() {
    setDataRoom({});
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/datarooms';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(DataRoom),
    };

    console.log(DataRoom); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.id != null) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
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
      <Header style={HeaderCustom} title={`ข้อมูลห้องพัก`}>
        <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10, marginRight:20 }}>{cookieName}</div>
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
              <div className={classes.paper}>หมายเลขห้องพัก</div>
            </Grid>
            <Grid item xs={9}>
            <TextField  
            value = {DataRoom.RoomNumber || ''}
            label="เลขห้องพัก" 
            name = "RoomNumber"
            variant="outlined" 
            onChange={handleChange}/>
           </Grid>
            
            <Grid item xs={3}>
              <div className={classes.paper}>สถานะห้องพัก</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกสถานะห้องพัก</InputLabel>
                <Select
                  name="StatusRoom"
                  value={DataRoom.StatusRoom || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {StatusRoom.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.statusName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>โปรโมชัน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกโปรโมชัน</InputLabel>
                <Select
                  name="Promotion"
                  value={DataRoom.Promotion  || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {Promotion.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.promotionName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ประเภทห้องพัก</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกประเภทห้องพัก</InputLabel>
                <Select
                   value={DataRoom.TypeRoom || ''} // (undefined || '') = ''
                   onChange={handleChange}
                   name="TypeRoom"
                >
                  {TypeRoom.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.typeName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ราคาห้องพัก</div>
            </Grid>
            <Grid item xs={9}>
            <TextField  
            value={DataRoom.Price || ''}
            label="ราคา" 
            name = "Price"
            variant="outlined" 
            onChange={handleChange}/>
           </Grid>
           <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกข้อมูล
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default DataRoom;
