import './App.css';
import {
    BrowserRouter as Router,
    Redirect,
    Route,
    Switch,
  } from 'react-router-dom'
import Container from './components/Container'

const App = () => {

    return (
        
        <div className='App'>
            <div className='Navbar'>
                <a href='/facemasks'>face masks</a>
                <a href='/gloves'>gloves</a>
                <a href='/beanies'>beanies</a>
            </div>
            <Router>
                <Switch>
                    <Route path="/:category">
                        <Container />
                    </Route>
                </Switch>
                <Route exact path="/">
                    <Redirect to="/facemasks"/>
                </Route>
            </Router>
        </div>
    );
};

export default App;
