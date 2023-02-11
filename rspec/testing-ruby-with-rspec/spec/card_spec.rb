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
  # load every time without lazy loading
  # before do 
  #   @card = Card.new('Ace', 'Spade')
  # end
  
  # test helper
  # def card 
  #   Card.new('Ace', 'Spade')
  # end

  # memoization with rspec 'let'
  # lazy loading, load when needed
  let(:card) { Card.new('Ace', 'Spade') }
  let(:x) { 1 + 1}
  let(:y) { x + 10}

  # let with bang, before example run ruby block below
  # let!(:card) { Card.new('Ace', 'Spade') }

  it 'testing rspec let' do 
    puts y
  end
  
  it 'has a rank that rank can change' do 
    expect(card.rank).to eq('Ace')

    card.rank = 'Queen'
    expect(card.rank).to eq('Queen')
  end

  it 'has a suit' do
    expect(card.suit).to eq('Spade')
  end

  # it 'has a custom error message' do 
  #   card.suit = "Nonsense"
  #   comparison = 'Spades'
  #   expect(card.suit).to eq(comparison), "I expected #{comparison} but I got #{card.suit} instead"
  # end
end

# matchers
# https://www.rubydoc.info/gems/rspec-expectations/RSpec/Matchers
