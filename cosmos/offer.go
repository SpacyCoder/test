package cosmos

type Offer struct {
	client  Client
	offerID string
}

type OfferDefinition struct {
	OfferVersion string `json:"offerVersion"`
	OfferType    string `json:"offerType"`
	Content      struct {
		OfferThroughput int `json:"offerThroughput"`
	} `json:"content"`
	Resource        string `json:"resource"`
	OfferResourceID string `json:"offerResourceId"`
	ID              string `json:"id"`
	Rid             string `json:"_rid"`
	Self            string `json:"_self"`
	Etag            string `json:"_etag"`
	Ts              int    `json:"_ts"`
}

type UserDefinition struct {
	Resource
	_persmissions string `json:"_persmissions,omitempty"`
}

type Offers struct {
	client Client
}

/*

func newOffer(client Client, offerID string) *Offer {
	client.path += "/offers/" + offerID
	client.rType = "offers"
	client.rLink = client.path
	offer := &Offer{
		client:  client,
		offerID: offerID,
	}

	return offer
}

func newOffers(client Client) *Offers {
	client.path += "/offers"
	client.rType = "offers"
	offers := &Offers{
		client: client,
	}

	return offers
}

func (u *Offers) Create(user *UserDefinition, opts ...CallOption) (*UserDefinition, error) {
	createdUser := &UserDefinition{}
	_, err := u.client.create(user, &createdUser, opts...)
	if err != nil {
		return nil, err
	}

	return createdUser, err
}

func (u *Offer) Replace(user *UserDefinition, opts ...CallOption) (*UserDefinition, error) {
	updatedUser := &UserDefinition{}
	_, err := u.client.replace(user, &updatedUser, opts...)
	if err != nil {
		return nil, err
	}

	return updatedUser, err
}

func (o *Offers) ReadAll(opts ...CallOption) ([]UserDefinition, error) {
	data := struct {
		Offers []UserDefinition `json:"offers,omitempty"`
		Count  int              `json:"_count,omitempty"`
	}{}

	_, err := o.client.read(&data, opts...)
	if err != nil {
		return nil, err
	}
	return data.Users, err
}

func (o *Offer) Delete(opts ...CallOption) (*Response, error) {
	return o.client.delete(opts...)
}

func (o *Offer) Read(opts ...CallOption) (*UserDefinition, error) {
	user := &UserDefinition{}
	_, err := o.client.read(user, opts...)
	return user, err
}
*/
