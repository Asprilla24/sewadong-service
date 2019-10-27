class CreateRoleTable < ActiveRecord::Migration[5.2]
  def change
      create_table :roles, id: false do |t|
        t.string :role_id, limit: 128, null: false, primary_key: true
        t.string :name, limit: 128, null: false
      end

      create_table :products, id: false do |t|
        t.string :product_id, limit: 128, null: false, primary_key: true
        t.string :name, limit: 128, null: false
        t.string :description, limit: 225, null: false
        t.string :image, limit: 225, null: false
        t.decimal :price, null: false
        t.string :payment_id, limit: 128, null: false
        t.string :category_id, limit: 128, null: false
        t.string :user_id, limit: 128, null: false
        t.datetime :updated_at, null: false, :default => Time.now
        t.datetime :created_at, null: false, :default => Time.now
      end

      create_table :categories, id: false do |t|
        t.string :category_id, limit: 128, null: false, primary_key: true
        t.string :name, limit: 128, null: false
      end

      execute "insert into roles (role_id, name) values ('1', 'Tenant'),('2', 'Regular User')"
  end
end
