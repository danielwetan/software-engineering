RSpec.describe '#even? method' do 
  # it 'should return true if number is even' do 
  # end

  # it 'should return false if number is odd' do 
  # end

  # nested describe
  # describe 'with even number' do 
  #   it 'should return true' do 
  #     expect(4.even?).to eq(true)
  #   end
  # end

  # grouping test with context
  context 'with even number' do 
    it 'should return true' do 
      expect(4.even?).to eq(true)
    end
  end

  context 'with odd number' do 
    it 'should return false' do 
      expect(5.even?).to eq(false)
    end
  end
end
