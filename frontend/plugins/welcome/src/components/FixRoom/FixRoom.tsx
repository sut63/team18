import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import {
  MuiPickersUtilsProvider,
  KeyboardTimePicker,
  KeyboardDatePicker,
} from '@material-ui/pickers';
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
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntDataRoom } from '../../api/models/EntDataRoom';
import { EntCustomer } from '../../api/models/EntCustomer';
import { EntReserveRoom } from '../../api/models/EntReserveRoom';
import { EntFurnitureDetail } from '../../api/models/EntFurnitureDetail';
import { Cookies } from '../../Cookie'
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

interface FixRoom {
  Room: number;
  FurnitureDetail: number;
  Customer: number;
  FixDetail: string;
  Tel: string;
  Facebook: string;
  // create_by: number;
}

const FixRoom: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();

  // ดึงคุกกี้
  var ck = new Cookies()
  var cookieName = ck.GetCookie()
  var cookieID = ck.GetID()

  //อันหลักสำหรับสร้าง การ reserve_room
  const [FixRoom, setFixRoom] = React.useState<Partial<FixRoom>>({});

  // DataRoom 
  const [DataRoom, setDataRoom] = React.useState<EntDataRoom[]>([])
  const [dids, setDIds] = React.useState<number>(0)
  const getDataRoom = async () => {
    const res = await api.listDataroom({ limit: 10, offset: 0 })
    setDataRoom(res)
  }

  // furnituredetail List
  const [FurnitureDetail, setFurnitureDetails] = React.useState<EntFurnitureDetail[]>([]);
  const getFurnitureDetail = async () => {
    const res = await api.listFurnitureDetailRoom({ limit: 10, offset: 0, id: dids });
    setFurnitureDetails(res);
  };

  //customer
  const [customer, setCustomer] = React.useState<EntCustomer>();
  const getCustomer = async () => {
    const res = await api.getCustomer({ id: Number(cookieID) });
    setCustomer(res);
  };

  //alert setting
  const [open, setOpen] = React.useState(false);
  const [fail, setFail] = React.useState(false);

   // สำหรับตรวยสอบความถูกต้อง
   const [FixDetailError, setFixDetailError] = React.useState('');
   const [PhoneNumberError, setPhoneNumberError] = React.useState('');
   const [FacebookError, setFacebookError] = React.useState('');
   const [errors, setError] = React.useState(String);
  
  //validate
  const validatePhoneNumber = (val: string) => {
    return val.match("[0]\\d{9}") && val.length <= 10;
  }

  const validateFacebook = (val: string) => {
    return val.length > 50 ? false : true;
  }

  // ฟังก์ชั่นสำหรับ validate รายละเอียดอุปกรณ์ที่ชำรุด
  const validateFixDetail = (val: string) => {
    return val.length > 50 ? false : true;
  }

  // checkPattern
  const checkPattern  = (id: string, value: string) => {
    switch(id) { 
      case 'Facebook':
        validateFacebook(value) ? setFacebookError('') : setFacebookError('ห้ามเกิน 50 ตัวอักษร');
        return;
      case 'Tel': 
        validatePhoneNumber(value) ? setPhoneNumberError('') : setPhoneNumberError('Ex 0850583300');
        return;
      case 'FixDetail':
        validateFixDetail(value) ? setFixDetailError('') : setFixDetailError('ห้ามเกิน 50 ตัวอักษร');
        return;
      default:
        return;
    }
  }

  //กำหนดข้อความ error
  const checkerror = (s :string) => {
    switch(s) {
      case 'fix_detail':
        setError("จำนวนตัวอักษรเกิน 50 ตัวอักษร")
        return;
      case 'phone_number':
        setError("รูปแบบเบอร์โทรศัพท์ไม่ถูกต้อง")
        return;
      case 'facebook':
        setError("จำนวนตัวอักษรเกิน 50 ตัวอักษร")
        return;
      default:
        setError("กรุณกรอกข้อมูลให้ครบถ้วน")
        return;
    }
  }; 
  // Lifecycle Hooks
  useEffect(() => {
    getCustomer();
    getDataRoom();
    getFurnitureDetail();
  }, []);

  useEffect(() => {
    getFurnitureDetail();
  }, [FixRoom.Room]);

  useEffect(() => {
    setFixRoom({ ...FixRoom, ['Customer']: customer?.id });
  }, [customer]);

  // set data to object reserve_room
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof FixRoom;
    const { value } = event.target;
    setFixRoom({ ...FixRoom, [name]: value });
    setDIds(event.target.value);
    const validateValue = value.toString() 
    checkPattern(name, validateValue)
    console.log(FixRoom);
  };

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }
    setFail(false);
    setOpen(false);
  };

  // clear input form
  function clear() {
    setFixRoom({});
    getCustomer();
  }

  // clear cookies
  function Clears() {
    ck.ClearCookie()
    window.location.reload(false)
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/fixrooms';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(FixRoom),
    };

    console.log(FixRoom); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          clear();
          setOpen(true);
        } else {
          checkerror(data.error.Name);
          setFail(true);
        }
      });
  }

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบแจ้งซ่อม`}>
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
              <div className={classes.paper}>ลูกค้า</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  disabled
                  name="Customer"
                  variant="outlined"
                  value={customer?.name}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>Room</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกห้อง</InputLabel>
                <Select
                  name="Room"
                  value={FixRoom.Room || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {DataRoom.map(item => {
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
              <div className={classes.paper}>อุปกรณ์ที่ชำรุด</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>อุปกรณ์ที่ชำรุด</InputLabel>
                <Select
                  name="FurnitureDetail"
                  value={FixRoom.FurnitureDetail || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {FurnitureDetail.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.edges?.furnitures?.furnitureName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>รายละเอียดอุปกรณ์ที่ชำรุด</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  error = {FixDetailError ? true : false}
                  helperText={FixDetailError}
                  name="FixDetail"
                  variant="outlined"
                  value={FixRoom.FixDetail || ''}
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
                  error = {PhoneNumberError ? true : false}
                  helperText={PhoneNumberError}
                  name="Tel"
                  label="เบอร์โทรศัพท์"
                  variant="outlined"
                  value={FixRoom.Tel || ''}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>Facebook</div>
            </Grid>
            <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  error = {FacebookError ? true : false}
                  helperText={FacebookError}
                  name="Facebook"
                  label="Facebook"
                  variant="outlined"
                  value={FixRoom.Facebook || ''}
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
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกการดู
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

export default FixRoom;
