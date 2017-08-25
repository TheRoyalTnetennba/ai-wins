import React from 'react';
import { Provider } from 'react-redux';
import { HashRouter, Route } from 'react-router-dom';

const App = ({ store }) => (
      <Provider store={store}>
        <HashRouter>
          <Switch>
            <ProtectedRoute exact path="/campaigns" component={CampaignListContainer} />
            <ProtectedRoute exact path="/campaigns/new" component={CampaignFormContainer} />
            <ProtectedRoute path="/campaigns/:campaignID" component={CampaignShowContainer} />
            <ProtectedRoute path="/search/:search" component={CampaignListContainer} />
            <ProtectedRoute path="/users/:userID" component={ProfileContainer} />
            <Route path="/session" component={AuthModalContainer} />
            <Route path="/" exact component={CampaignIndexContainer} />
          </Switch>
        </HashRouter>
      </Provider>
);
export default App;
