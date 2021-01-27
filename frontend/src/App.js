import 'bootstrap/dist/css/bootstrap.css';
import './App.css';

function App() {
  const notImplemented = () => {
    alert('Not implemented');
  }
  return (
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          <div class="hislist">
            <h1>Http Info Service</h1>
            <div class="row">
              <div class="col-md-10">
                <input type="text" class="form-control add-his" placeholder="Add his" />
              </div>
              <div class="col-md-2">
                <button class="btn btn-success" onClick={() => notImplemented()}>Add & Save</button>
              </div>
            </div>

            <hr />
            <ul class="list-unstyled info-list">
              <li>
                <span class="align-middle">Чурт санна лаьтта со гIумгIазин арахь</span>
              </li>
              <li>
                <span class="align-middle">Иштта хIуманаш ду</span>
              </li>
              <li>
              <span class="align-middle">Кхин а информаци</span>
              </li>
            </ul>
            <div class="his-footer">
              <strong>
                <span class="count-hiss">3</span>
              </strong>{' '}
              infos total
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
