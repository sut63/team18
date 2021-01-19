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
import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';

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
  RoomDetail: string;
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
  DataRoom.Price = Number(DataRoom.Price)
  const [RoomnumberError, setRoomnumberError] = React.useState('');
  const [PriceError, setPriceError] = React.useState('');
  const [RoomDetailError, setRoomdetailError] = React.useState('');
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


  const validateRoomnumber = (val: string) => {
    return val.match("[A-Z]\\d{3}");
  }

  
  const validatePrice = (val: Number) => {
    return val > 0 ? true : false;
  }

  // ฟังก์ชั่นสำหรับ validate รหัสนักศึกษา
  const validateRoomdetail = (val: string) => {
    return val.length > 70 ? false : true;
  }
//checkpatten
  const checkPattern  = (id: string, value: string) => {
    switch(id) {
      case 'RoomNumber':
        validateRoomnumber(value) ? setRoomnumberError('') : setRoomnumberError('หมายเลขห้องขึ้นต้นด้วย A-Z เเละ ตัวเลข 3 ตัว');
        return;
      case 'Price':
        validatePrice(Number(value)) ? setPriceError('') : setPriceError('เป็นจำนวนเต็มบวกเท่านั้น');
        return;
      case 'RoomDetail':
        validateRoomdetail(value) ? setRoomdetailError('') : setRoomdetailError('ห้ามเกิน 70 ตัวอักษร')
        return;
      default:
        return;
    }
  }

  // alert setting
  const [open, setOpen] = React.useState(false);
  const [fail, setFail] = React.useState(false);
  const [errors, setError] = React.useState(String);
  //close alert 
  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }
    setFail(false);
    setOpen(false);
  };

  const checkerror = (s :string) => {
    switch(s) {
      case 'roomnumber':
        setError("รูปแบบหมายเลขห้องพักไม่ถูกต้อง")
        return;
      case 'price':
        setError("ราคาไม่ถูกต้อง")
        return;
      case 'roomdetail':
        setError("รายละเอียดมีความยาวมากเกินไป")
        return;
      default:
        return;
    }
  };
  
  
  // Lifecycle Hooks
  useEffect(() => {
    getStatusRoom();
    getPromotion();
    getTypeRoom();
  }, []);

  // set data to object DataRoom
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof DataRoom;
    const { value } = event.target;
    setDataRoom({ ...DataRoom, [name]: value });
    const validateValue = value.toString() 
    checkPattern(name, validateValue)
    console.log(DataRoom);
  };
  


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
        console.log(data.status);
        if (data.status == true) {
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

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ข้อมูลห้องพัก`}>
        <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
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
              <div className={classes.paper}>หมายเลขห้องพัก</div>
            </Grid>
            <Grid item xs={9}>
              <TextField
                error = {RoomnumberError ? true : false}
                value={DataRoom.RoomNumber || ''}
                helperText={RoomnumberError}
                label="เลขห้องพัก"
                name="RoomNumber"
                variant="outlined"
                onChange={handleChange} />
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
                  value={DataRoom.Promotion || ''} // (undefined || '') = ''
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
                error = {PriceError ? true : false}
                helperText={PriceError}
                type={"number"}
                value={DataRoom.Price || ''}
                label="ราคา"
                name="Price"
                variant="outlined"
                onChange={handleChange} />
            </Grid>
           

            <Grid item xs={3}>
              <div className={classes.paper}>รายละเอียด</div>
            </Grid>
            <Grid item xs={9}>
              <TextField
                error = {RoomDetailError ? true : false}
                helperText={RoomDetailError}
                value={DataRoom.RoomDetail || ''}
                label="รายละเอียด"
                name="RoomDetail"
                variant="outlined"
                onChange={handleChange} />
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

          <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="success">
              This is a success message!
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

export default DataRoom;
