
import React from 'react';
import { AppBar,Layout } from 'react-admin';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';

 const submitHandler = (e:any) => {
    e.preventDefault()
    fetch('http://attack-seaman.com:6868/attackPatterns/save',{
        method: 'POST',
    }).then(response => {
            console.log(response)
        })
        .catch(error =>{
            console.log(error)
        })
}

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
    return (
      <AppBar {...props} color='secondary' >
        <Typography
          variant="h6"
          color="inherit"
          className={classes.title}
        >Mitre Attack 编辑</Typography>
        <Typography
          variant="h6"
          color="inherit"
          className={classes.title}
          id="react-admin-title"
        />
        <Button style= {
          {
            color: '#FFFFFF',
            fontWeight: 500,
            fontSize: 14,
          }} 
        onClick={submitHandler}>刷新layer</Button>
      </AppBar >
    );
  };
  
//  export default CustomAppBar;
export const CustomLayout = (props:any) => <Layout {...props} appBar={CustomAppBar} />;

// export default CustomLayout;