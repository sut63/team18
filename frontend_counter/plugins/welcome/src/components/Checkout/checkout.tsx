import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import {  Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import {Cookies} from '../../Cookie'
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
  Typography
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
// me
import { EntCheckIn } from '../../api/models/EntCheckIn'; // import interface checkin
import { EntStatus } from '../../api/models/EntStatus'; // import interface Status
import { EntCounterStaff } from '../../api/models/EntCounterStaff'; // import interface CounterStaff
import { EntStatusOpinion } from '../../api/models/EntStatusOpinion'; // import interface status opinion
import { EntReserveRoom } from '../../api/models/EntReserveRoom'; // import interface status ReserveRoom

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
    width: 500,
  },
}));

interface CheckOut {
    CheckinsID: number;
    StatussID: number;
    CounterstaffsID: number;
    Identitycard:    string;
	  StatusopinionID: number;
	  Comment:         string;
	  Price:           number;
}

const checkout: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();

  // Cookie
  var ck = new Cookies()
  var cookieName = ck.GetCookie()
  var cookieID = ck.GetID()

  const [CheckOut, setCheckout] = React.useState<Partial<CheckOut>>({});
  CheckOut.Price = Number(CheckOut.Price)
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



  // get chackin
  const [Checkin, setCheckin] = React.useState<EntCheckIn[]>([]); 
  const getcheckIn = async () => {
    const res = await api.listGetCheckInStatus();
    setCheckin(res);
  };

  //get status pay
  const [Statusid, setStatusid] = React.useState(1);
  const [Status, setStatus] =   React.useState<EntStatus[]>([]);
  const getstatus = async () => {
    const res = await api.listStatus({ limit: 10, offset: 0 });
    setStatus(res);
  };

  // get couterstaff
  const [counter, setCounter] = React.useState<EntCounterStaff>()
  const getCounterStaff = async () => {
    const res = await api.getCounterStaff({ id: Number(cookieID) })
    setCounter(res)
  }
  
  // get opinion
  const [opinion, setopinion] = React.useState<EntStatusOpinion[]>([]);
  const getstatusopinion = async () => {
    const res = await api.listStatusopinion({ limit: 10, offset: 0 });
    setopinion(res);
  };
  // Lifecycle Hooks
  useEffect(() => {

    getstatusopinion(); 
    getcheckIn();
    getstatus();
    getCounterStaff();
  }, [CheckOut.CheckinsID,CheckOut.CounterstaffsID,CheckOut.StatussID,CheckOut.StatusopinionID]);



  useEffect(() => {
    setCheckout({ ...CheckOut, ['CounterstaffsID']: counter?.id })
  }, [counter]);

  
  // สำหรับตรวจสอบความถูกต้อง
  const [Identitycard, setIdentitycardError] = React.useState('');
  const [Comment, setCommentError] = React.useState('');
  const [Price, setPriceError] = React.useState('');

  

  //valid functions  
  const validateIdentitycard = (val: string) => {
    return val.length == 13 ? true:false;
  }

  const validatePrice = (val: Number) => {
    return val > 0 ? true : false;
  }

  const validateComment = (val: string) => {
    return val.length > 70 ? false : true;
  }

  // checkPattern
  const checkPattern  = (id: string, value: string) => {
    switch(id) { 
      case 'Comment':
        validateComment(value) ? setCommentError('') : setCommentError('ความเห็นยาวเกินขนาด 70 ตัวอักษร');
        return;
      case 'Price': 
      validatePrice(Number(value)) ? setPriceError('') : setPriceError('จำนวนเงินไม่ถูกต้อง');
        return;
      case 'Identitycard':
        validateIdentitycard(value) ? setIdentitycardError('') : setIdentitycardError('ตัวเลขไม่ครบ 13 หลัก')
        return;
      default:
        return;
    }
  }


// func checkerror
const [errors, setError] = React.useState(String);
const checkerror = (s :string) => {
  switch(s) {
    case 'identity_card':
      setError("เลขประจำตัวไม่ครบ 13 หลัก") 
      return;
    case 'price':
      setError("ตัวเลขไม่ถูกต้อง")
      return;
    case 'comment':
      setError("ข้อความยาวเกิน 70 ตัวอักษร")
      return;
    default:
      setError("บันทึกไม่สำเร็จ")
      return;
  }
};

  // set data to object DataRoom
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof CheckOut;
    const { value } = event.target;
    setCheckout({ ...CheckOut, [name]: value });
    
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    if (name == "CheckinsID"){
      getReserveroom(value)
    } 
  };

  
  //get reserveroom
  const [idReserveroom, setidReserveroom] = React.useState<Number>(0); 
  const [Reserveroom, setReserveroom] = React.useState<EntReserveRoom>(); 
  const getReserveroom = async (id:number) => {
    const res = await api.getGetCheckout2({ id:id})
    console.log("res = ==== ");
    console.log(res);
    setReserveroom(res)
  }
  useEffect(() => {

  },[])
  

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/checkouts';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(CheckOut),
    };

    console.log(CheckOut); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
    .then(response => response.json())
    .then(data => {
      console.log(data.status);
      console.log(data);
      if (data.status == true) {
        clear();
        setOpen(true);
        
      } else {
        clear();
        checkerror(data.error.Name);
        setFail(true);
        console.log(data);
       
      }
    });
  }

// clear input form
  function clear() {
    setCheckout({});
    getCounterStaff();
  }

  function Clears() {
    ck.ClearCookie()
    window.location.reload(false)
  }
  function reload() {
    window.location.reload(false)
  }
  function wait(ms:number){
    var start = new Date().getTime();
    var end = start;
    while(end < start + ms) {
      end = new Date().getTime();
   }
 }

 const timer = setTimeout(()=>{
  window.location.reload(false);
},5000) //  5000 = 5 second 

  return (
    <Page theme={pageTheme.home}>
  
      <Header style={HeaderCustom} title={`ระบบ check out`}>
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
            
            
            <Grid item xs={3}>
              <div className={classes.paper}>check in</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือก ID checkin</InputLabel>
                <Select
                  name="CheckinsID"
                  value={CheckOut.CheckinsID  || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {Checkin.map(item => {
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
              <div className={classes.paper}>counter staff</div>
            </Grid>
            <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  disabled
                  name="CounterstaffsID"
                  variant="outlined"
                  value={counter?.name}
                />
              </FormControl>
            </Grid>
            <Grid item xs={12}>
            <Typography variant="h5"  color="primary" >จำนวนเงินที่ต้องจ่าย {Reserveroom?.netPrice}</Typography>
            </Grid>
            <Grid item xs={3}>
            <div className={classes.paper}>จำนวนเงินที่จ่าย</div>
            </Grid>
            <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
            <TextField
                error = {Price ? true : false}
                helperText={Price}
                label="ใส่จำนวนเงิน"
                type={"number"}
                name="Price"
                value={CheckOut.Price || ''}
                variant="outlined"
                onChange={handleChange} />
                </FormControl>
           </Grid>
           <Grid item xs={3}>
              <div className={classes.paper}>status</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกสถานะการชำระเงิน</InputLabel>
                <Select
                  name="StatussID"
                  value={CheckOut.StatussID  || 1 } // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {Status.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.description}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
           <Grid item xs={3}>
            <div className={classes.paper}>เลขบัตรประชาชน</div>
            </Grid>
            <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
            <TextField
                error = {Identitycard ? true : false}
                helperText={Identitycard}
                label="ใส่เลขบัตรประชาชน"
                name="Identitycard"
                variant="outlined"
                onChange={handleChange} />
                </FormControl>
           </Grid>

           <Grid item xs={3}>
              <div className={classes.paper}>ความพึงพอใจลูกค้า</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกความพึงหอใจ</InputLabel>
                <Select
                  name="StatusopinionID"
                  value={CheckOut.StatusopinionID  || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {opinion.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.opinion}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>        

           <Grid item xs={3}>
            <div className={classes.paper}>ความเห็นลูกค้า</div>
            </Grid>
            <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
            <TextField
              multiline
              rows={4}
              error = {Comment ? true : false}
                helperText={ Comment}
                label="ใส่ความเห็น"
                name="Comment"
                variant="outlined"
                onChange={handleChange} />
                 </FormControl>
           </Grid>


           <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={() => {
                  save();
                }}
              >
                check out
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

export default checkout;
