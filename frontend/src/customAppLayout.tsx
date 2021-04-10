import React from 'react';
import { AppBar, Layout, FormTab, FileInput } from 'react-admin';
import { Typography,Button,Fade,Snackbar } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles({
  title: {
    flex: 1,
    textOverflow: 'ellipsis',
    whiteSpace: 'nowrap',
    overflow: 'hidden',
    marginLeft: -10
  },
  spacer: {
    flex: 1,
  },
  logo: {
    maxWidth: "40px",
    marginLeft: -35
  },
});

const CustomAppBar = (props:any) => {
  const classes = useStyles();
  const [state, setState] = React.useState({
    open: false,
    Transition: Fade,
    msg: '',
  });

  const handleFileSelected = (e: React.ChangeEvent<HTMLInputElement>): void => {
    const files = e.target.files;
    console.log("files:", files);

    const formData = new FormData();

    if (files) {
      formData.append("File", files[0]);

      fetch("http://attack-seaman.com:6868/upload?", {
        method: "POST",
        body: formData,
      })
        .then((response) => response.json())
        .then((result) => {
          console.log("Success:", result);
        })
        .catch((error) => {
          console.error("Error:", error);
        });
    }
  };

  const submitHandler = (e:any) => {
    e.preventDefault()
    setState({
      open: true,
      Transition: Fade,
      msg: "刷新中......"
    });
    fetch('http://attack-seaman.com:6868/attackPatterns/save',{
      method: 'POST',
    }).then(response => {
      console.log(response)
      setState({
        open: true,
        Transition: Fade,
        msg: "刷新成功！",
      });
    })
    .catch(error =>{
      console.log(error)
      setState({
        open: true,
        Transition: Fade,
        msg: "刷新失败！",
      });
    })
  }

  const handleClose = () => {
    setState({
      ...state,
      open: false,
    });
  };

  return (
    <AppBar {...props} color="secondary">
      <Typography variant="h6" color="inherit" className={classes.title}>
        Mitre Attack 编辑
      </Typography>
      <Typography
        variant="h6"
        color="inherit"
        className={classes.title}
        id="react-admin-title"
      />
      <Button variant="contained" component="label">
        Upload File
        <input type="file" id="fileinput" accept='.json' hidden onChange={handleFileSelected}/>
      </Button>
      <Button
        style={{
          color: "#FFFFFF",
          fontWeight: 500,
          fontSize: 14,
        }}
        onClick={submitHandler}
      >
        刷新layer
      </Button>
      <Snackbar
        open={state.open}
        onClose={handleClose}
        TransitionComponent={state.Transition}
        message={state.msg}
        autoHideDuration={6000}
      />
    </AppBar>
  );
};

//  export default CustomAppBar;
export const CustomLayout = (props:any) => <Layout {...props} appBar={CustomAppBar} />;

// export default CustomLayout;