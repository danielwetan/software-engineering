RSpec.describe 'nested hooks' do 
  before(:context) do 
    puts 'OUTER Before context'
  end

  before(:example) do 
    puts 'OUTER Before example'
  end

  it 'does basic math' do 
    expect(1 + 1).to eq(2)
  end

  # nested hooks
  context 'with condition A' do 
    before(:context) do 
      puts 'INNER Before context'
    end
  
    before(:example) do 
      puts 'INNER Before example'
    end

    it 'does some more basic math' do 
      expect(1 + 1).to eq(2)
    end

    it 'does substraction at all' do 
      expect(5 - 2).to eq(3)
    end
  end
end

