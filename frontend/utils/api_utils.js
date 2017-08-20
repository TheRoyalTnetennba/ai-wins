

export const login = credentials => (
  fetch({
    url: '/api/session',
    method: 'POST',
    data: credentials,
  })
);

export const signup = info => (
  fetch({
    url: 'api/users',
    method: 'POST',
    data: info,
  })
);

export const logout = () => (
  fetch({
    url: 'api/session',
    method: 'DELETE',
  })
);

export const guestLogin = () => (
  fetch({
    url: 'api/session/guest_user',
    method: 'GET',
  })
);

export const fetchAllCampaigns = () => (
  fetch({
    url: '/api/campaigns',
    method: 'GET',
  })
);

export const fetchCampaign = id => (
  fetch({
    url: `/api/campaigns/${id}`,
    method: 'GET',
  })
);

export const fetchAllCategories = () => (
  fetch({
    url: '/api/categories',
    method: 'GET',
  })
);

export const searchCampaigns = string => (
  fetch({
    url: `/api/campaigns/search/${string}`,
    method: 'GET',
  })
);

export const createCampaign = campaign => (
  fetch({
    url: 'api/campaigns',
    method: 'POST',
    data: campaign,
  })
);

export const createPerk = perk => (
  fetch({
    url: 'api/perks',
    method: 'POST',
    data: perk,
  })
);

export const sendContribution = contribution => (
  fetch({
    url: 'api/contributions',
    method: 'POST',
    data: contribution,
  })
);

export const fetchUser = user => (
  fetch({
    url: `api/users/${user}`,
    method: 'GET',
  })
);

export const submitContact = contact => (
  fetch({
    url: 'api/emails/',
    method: 'POST',
    data: contact,
  })
);
