class CreateUserTable < ActiveRecord::Migration[5.2]
  def change
    create_table :users, id: false do |t|
      t.string :user_id, limit: 128, null: false, primary_key: true
      t.string :email, limit: 128, null: false, unique: true
      t.string :username, limit: 128, null: false, unique: true
      t.string :password, limit: 128, null: false
      t.string :role_id, null:false
      t.string :gender, null:true
      t.string :phone_number, null:false
      t.string :address, null:false
      t.string :image, null:true
      t.datetime :updated_at, null: false, :default => Time.now
      t.datetime :created_at, null: false, :default => Time.now
    end
  end
end
