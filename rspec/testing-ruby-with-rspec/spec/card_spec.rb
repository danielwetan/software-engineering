class Card 
  attr_accessor :rank, :suit

  def initialize(rank, suit)
    @rank = rank 
    @suit = suit
  end
end

# https://relishapp.com/rspec/rspec-core/v/3-8/docs/example-groups/basic-structure-describe-it
RSpec.describe Card do
  # run ruby block before each test
  # before do 
  #   @card = Card.new('Ace', 'Spade')
  # end

  def card 
    Card.new('Ace', 'Spade')
  end
  
  it 'has a rank that rank can change' do 
    expect(card.rank).to eq('Ace')

    card.rank = 'Queen'
    expect(card.rank).to eq('Queen')
  end

  it 'has a suit' do
    expect(card.suit).to eq('Spade')
  end
end

# matchers
# https://www.rubydoc.info/gems/rspec-expectations/RSpec/Matchers
